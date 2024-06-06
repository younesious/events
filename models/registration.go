package models

import (
	"database/sql"
	"errors"

	"github.com/younesious/events/db"
)

func RegisterForEvent(eventID, userID int64) error {
	db := db.GetDB()

	var event Event
	err := db.QueryRow("SELECT id FROM events WHERE id = ?", eventID).Scan(&event.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("event not found")
		}
		return err
	}

	_, err = db.Exec("INSERT INTO registrations (event_id, user_id) VALUES (?, ?)", eventID, userID)
	if err != nil {
		return err
	}

	return nil
}

func CancelRegistration(eventID, userID int64) error {
	db := db.GetDB()

	var registrationID int64
	err := db.QueryRow("SELECT id FROM registrations WHERE event_id = ? AND user_id = ?", eventID, userID).Scan(&registrationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("registration not found")
		}
		return err
	}

	_, err = db.Exec("DELETE FROM registrations WHERE id = ?", registrationID)
	if err != nil {
		return err
	}

	return nil
}
