package events

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aerzz23/visadiscordbot/api/config"
	"github.com/boltdb/bolt"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)

const (
	eventBucketKey = "events"
)

// Event is a struct which aids with JSON marshalling
type Event struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	EventDate time.Time `json:"date"`
	EventTime time.Time `json:"eventTime"`
	EventName string    `json:"eventName"`
	EventInfo string    `json:"eventInfo"`
	Alert     bool      `json:"alert"`
	AlertTime int       `json:"alertTime"`
}

// New makes a new Event type, used for JSON marshalling.
func New(eventDate string, eventTime string, eventName string, eventInfo string, alert bool, alertTime int) (*Event, error) {
	createdAt := time.Now()
	eventDateParsed, err := time.Parse("20060102", eventDate)
	if err != nil {
		log.Fatalf("Error parsing new event date: %s, error: %v", eventDate, err)
		return nil, err
	}
	eventTimeParsed, err := time.Parse("15:04", eventTime)
	if err != nil {
		log.Fatalf("Error parsing new event time: %s, error: %v", eventDate, err)
		return nil, err
	}

	return &Event{
		CreatedAt: createdAt,
		EventDate: eventDateParsed,
		EventTime: eventTimeParsed,
		EventName: eventName,
		EventInfo: eventInfo,
		Alert:     alert,
		AlertTime: alertTime,
	}, nil
}

// Create will add a new event to BoltDB.
// Will use attributes of Event instance to insert.
// Creates bucket for Events if not already there.
func (e *Event) Create(cfg *config.BotConfig, db *bolt.DB) error {
	// serialize event ready for easy insert into BoltDB
	eventSerialized, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("Error serializing event: %v", err)
		return err
	}

	// make bucket if not there and add event to db
	return db.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists([]byte(cfg.DB.Buckets[eventBucketKey]))

		if err != nil {
			log.Fatalf("Error creating bucket: %v", err)
			return err
		}

		key, err := b.NextSequence()

		if err != nil {
			log.Fatalf("Error getting next sequence in DB: %v", err)
			return err
		}

		keyBuffer := new(bytes.Buffer)

		err = binary.Write(keyBuffer, binary.LittleEndian, key)

		if err != nil {
			log.Fatalf("Error converting sequence db key into bytes. Error: %v", err)
			return err
		}

		err = b.Put(keyBuffer.Bytes(), eventSerialized)

		if err != nil {
			log.Fatalf("Error inserting new event into db. Error: %v", err)
			return err
		}

		return nil
	})
}

// Map takes in a slice of Events and will create a new slice of given type using the provided function.
func Map(inputs []Event, fn func(Event) interface{}) interface{} {
	inputsMapped := make([]interface{}, len(inputs))

	for i, v := range inputs {
		inputsMapped[i] = fn(v)
	}

	return inputsMapped
}

// FormatAsTable returns HTML table with events formatted.
func FormatAsTable(events []Event) string {
	eventsMapped := Map(events, func(e Event) interface{} {
		stringSlice := make([]string, 4)
		stringSlice[0] = e.EventDate.Format("2006-01-02")
		stringSlice[1] = e.EventName
		stringSlice[2] = e.EventTime.Format("15:04")
		stringSlice[3] = e.EventInfo
		return stringSlice
	})

	mappedEventsAsInterfaceSlices := eventsMapped.([]interface{})

	mappedEventsAsStringSlices := make([][]string, len(mappedEventsAsInterfaceSlices))

	for i, v := range mappedEventsAsInterfaceSlices {
		mappedEventsAsStringSlices[i] = v.([]string)
	}

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)

	table.SetHeader([]string{"Date", "Name", "Time", "Info"})
	table.SetBorder(true)
	table.AppendBulk(mappedEventsAsStringSlices)
	table.Render()

	return tableString.String()
}

// GetAll returns all Discord events in DB.
func GetAll(cfg *config.BotConfig, db *bolt.DB) ([]Event, error) {
	upcomingEvents := []Event{}
	err := db.View(func(tx *bolt.Tx) error {
		// TODO fix error when bucket is nil
		b := tx.Bucket([]byte(cfg.DB.Buckets[eventBucketKey]))
		err := b.ForEach(func(k, v []byte) error {
			var event Event
			err := yaml.Unmarshal(v, &event)
			if err != nil {
				log.Println(fmt.Sprintf("Error whilst unmarshalling event row from DB. Error: %v", err))
				return err
			}
			upcomingEvents = append(upcomingEvents, event)
			return nil
		})

		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return upcomingEvents, err
	}

	return upcomingEvents, nil
}

// TODO add upcoming events (so filtered by date/time in future)
// GetAllUpcoming returns all upcoming Discord events in DB.
// func GetAllUpcoming(db *bolt.DB) error {

// }
