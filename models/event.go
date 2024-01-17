package models

import (
	"time"

	"github.com/bhattrajat/go-events-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

func (e *Event) Save() error {
	stmt, err := db.DB.Prepare(`INSERT INTO events (name, description, location, dateTime, userId) 
		values(?, ?, ?, ?, ?)`)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, 1)
	defer stmt.Close()
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	getAllEventsQuery := `SELECT * FROM events`
	rows, err := db.DB.Query(getAllEventsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	getAllEventsQuery := `SELECT * FROM events where id=?`

	row := db.DB.QueryRow(getAllEventsQuery, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
