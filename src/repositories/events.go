package repositories

import (
	"PingChef/src/domain"
	"database/sql"
)

type event struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *event {
	return &event{db: db}
}

func (e event) Create(event domain.Event) error {
	statement, err := e.db.Prepare("INSERT INTO events (id, url, status, response_time) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(
		event.ID,
		event.URL,
		event.Status,
		event.ResponseTime,
	); err != nil {
		return err
	}

	return nil
}
