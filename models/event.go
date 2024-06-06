package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/younesious/events/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) CreateEvent() error {
	db := db.GetDB()
	query := `INSERT INTO events (name, description, location, date_time, user_id) 
                  VALUES (?, ?, ?, ?, ?)`

	res, err := db.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = id

	return nil
}

func GetAllEvents() ([]Event, error) {
	db := db.GetDB()
	query := `SELECT * FROM events`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func GetEvent(id int64) (*Event, error) {
	db := db.GetDB()

	query := `SELECT * FROM events WHERE id = ?`

	var event Event
	err := db.QueryRow(query, id).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("event not found!")
		}
		return nil, err
	}

	return &event, nil
}

func (e *Event) UpdateEvent() error {
	db := db.GetDB()
	query := `UPDATE events
	SET name = ?, description = ?, location = ?, date_time = ?
	WHERE id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = db.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteEvent(id int64) error {
	db := db.GetDB()

	query := `DELETE FROM events WHERE id = ?`

	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
