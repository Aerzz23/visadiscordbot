package events

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
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
func (e *Event) Create(db *bolt.DB) error {
	// serialize event ready for easy insert into BoltDB
	eventSerialized, err := json.Marshal(e)
	if err != nil {
		log.Fatalf("Error serializing event: %v", err)
		return err
	}

	// make bucket if not there and add event to db
	return db.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists([]byte("Events"))

		if err != nil {
			log.Fatalf("Error creating bucket: %v", err)
			return err
		}

		key, err := b.NextSequence()

		if err != nil {
			log.Fatalf("Error getting next sequence in DB: %v", err)
			return err
		}

		keyAsInt := int(key)

		keyBuffer := new(bytes.Buffer)

		err = binary.Write(keyBuffer, binary.LittleEndian, keyAsInt)

		err = b.Put(keyBuffer.Bytes(), eventSerialized)

		return nil
	})
}
