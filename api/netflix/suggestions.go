package netflix

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
	"time"
)

const (
	eventBucketKey = "Suggestions"
)

// Suggestion is a struct which aids with JSON marshalling.
type Suggestion struct {
	Title     string    `json:"title"`
	Users     []string  `json:"users"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// New makes a new Suggestion type, used for JSON marshalling.
func New(title, author string) *Suggestion {
	return &Suggestion{
		Title:     title,
		Users:     []string{author},
		UpdatedAt: time.Now(),
	}
}

// Create will add a new event to BoltDB.
// Will use attributes of Suggestion instance to insert.
// Creates bucket for Suggestions if not already there.
func (s *Suggestion) Create(db *bolt.DB, author string) error {

	err := checkBucket(db)
	if err != nil {
		log.Fatalf("Error creating bucket: %v", err)
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		newSuggestion := s
		b := tx.Bucket([]byte(eventBucketKey))
		v := b.Get([]byte(s.Title))
		// if title already there
		if v != nil {
			oldSuggestion := &Suggestion{}
			err = json.Unmarshal(v, oldSuggestion)
			if err != nil {
				log.Fatalf("Error unmarshalling Suggestion: %v", err)
				return err
			}
			if notRepeat(s.Users[0], oldSuggestion.Users) {
				newUsers := oldSuggestion.Users
				newUsers = append(newUsers, s.Users[0])
				newSuggestion = oldSuggestion
				newSuggestion.Users = newUsers
			}
		}
		newSuggestion.UpdatedAt = time.Now()
		suggestionSerialised, err := json.Marshal(newSuggestion)
		if err != nil {
			log.Fatalf("Error serializing Suggestion: %v", err)
			return err
		}

		err = b.Put([]byte(newSuggestion.Title), suggestionSerialised)
		if err != nil {
			log.Fatalf("Error setting key: %v value: %v error: %v", []byte(s.Title), suggestionSerialised, err)
			return err
		}
		return nil
	})
}

func checkBucket(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {

		_, err := tx.CreateBucketIfNotExists([]byte("Suggestions"))

		if err != nil {
			log.Fatalf("Error creating bucket: %v", err)
			return err
		}
		return nil
	})
}

func notRepeat(newUser string, currentUsers []string) bool {
	for _, usr := range currentUsers {
		if usr == newUser {
			return false
		}
	}
	return true
}