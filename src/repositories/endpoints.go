package repositories

import (
	"PingChef/src/domain"
	"database/sql"

	_ "modernc.org/sqlite"
)

type endpoint struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *endpoint {
	return &endpoint{db: db}
}

func (e endpoint) Create(endpoint *domain.EndPoint) error {
	statement, err := e.db.Prepare("INSERT INTO endpoints (id, name, url, interval) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(
		endpoint.ID,
		endpoint.Name,
		endpoint.URL,
		endpoint.Interval,
	); err != nil {
		return err
	}

	return nil
}
