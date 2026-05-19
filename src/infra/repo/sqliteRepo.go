package repo

import (
	"PingChef/src/modules/endpoints"
	"database/sql"
)

type SQLiteRepo struct {
	db *sql.DB
}

func NewSQLiteEndpointRepo(db *sql.DB) *SQLiteRepo {
	return &SQLiteRepo{db: db}
}

func (r *SQLiteRepo) SaveEndpoint(endpoint endpoints.Endpoint) error {
	return nil
}

func (r *SQLiteRepo) SaveEvent(event endpoints.Event, endpointID string) error {
	return nil
}

func (r *SQLiteRepo) FindByID(id string) (endpoints.Endpoint, error) {
	return endpoints.Endpoint{}, nil
}

func (r *SQLiteRepo) FindAllEndpoints() ([]endpoints.Endpoint, error) {
	return []endpoints.Endpoint{}, nil
}

func (r *SQLiteRepo) FindAllEvents() ([]endpoints.Event, error) {
	return []endpoints.Event{}, nil
}

func (r *SQLiteRepo) DeleteEndpoint(id string) error {
	return nil
}

func (r *SQLiteRepo) UpdateEndpoint(endpoint endpoints.Endpoint) error {
	return nil
}
