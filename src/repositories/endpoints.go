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

func (e endpoint) List() ([]domain.EndPoint, error) {
	rows, err := e.db.Query("SELECT id, name, url, interval FROM endpoints")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var endpoints []domain.EndPoint
	for rows.Next() {
		var endpoint domain.EndPoint
		if err := rows.Scan(&endpoint.ID, &endpoint.Name, &endpoint.URL, &endpoint.Interval); err != nil {
			return nil, err
		}
		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}
