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
	statement, err := r.db.Prepare("INSERT INTO endpoints (id, name, url, interval) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(endpoint.ID, endpoint.Name, endpoint.URL, endpoint.Interval)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepo) SaveEvent(event endpoints.Event, endpointID string) error {
	statement, err := r.db.Prepare("INSERT INTO events (id, endpoint_id, status, status_code, response_time) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.ID, endpointID, event.Status, event.StatusCode, event.ResponseTime)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepo) FindByID(id string) (endpoints.Endpoint, error) {
	statement, err := r.db.Prepare("SELECT id, name, url, interval FROM endpoints WHERE id = ?")
	if err != nil {
		return endpoints.Endpoint{}, err
	}
	defer statement.Close()

	row := statement.QueryRow(id)
	var endpoint endpoints.Endpoint
	if err := row.Scan(&endpoint.ID, &endpoint.Name, &endpoint.URL, &endpoint.Interval); err != nil {
		return endpoints.Endpoint{}, err
	}

	eventStatement, err := r.db.Prepare("SELECT id, status, status_code, response_time FROM events WHERE endpoint_id = ?")
	if err != nil {
		return endpoints.Endpoint{}, err
	}
	defer eventStatement.Close()

	eventRows, err := eventStatement.Query(id)
	if err != nil {
		return endpoints.Endpoint{}, err
	}
	defer eventRows.Close()

	for eventRows.Next() {
		var event endpoints.Event
		if err := eventRows.Scan(&event.ID, &event.Status, &event.StatusCode, &event.ResponseTime); err != nil {
			return endpoints.Endpoint{}, err
		}
		endpoint.AddEvent(event)
	}

	return endpoint, nil
}

func (r *SQLiteRepo) FindAllEndpoints() ([]endpoints.Endpoint, error) {
	statement, err := r.db.Prepare("SELECT id, name, url, interval FROM endpoints")
	if err != nil {
		return []endpoints.Endpoint{}, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return []endpoints.Endpoint{}, err
	}
	defer rows.Close()

	var endpointsList []endpoints.Endpoint
	for rows.Next() {
		var endpoint endpoints.Endpoint
		if err := rows.Scan(&endpoint.ID, &endpoint.Name, &endpoint.URL, &endpoint.Interval); err != nil {
			return []endpoints.Endpoint{}, err
		}
		endpointsList = append(endpointsList, endpoint)
	}

	return endpointsList, nil
}

func (r *SQLiteRepo) FindAllEvents() ([]endpoints.Event, error) {
	statement, err := r.db.Prepare("SELECT id, status, status_code, response_time FROM events")
	if err != nil {
		return []endpoints.Event{}, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return []endpoints.Event{}, err
	}
	defer rows.Close()

	var eventsList []endpoints.Event
	for rows.Next() {
		var event endpoints.Event
		if err := rows.Scan(&event.ID, &event.Status, &event.StatusCode, &event.ResponseTime); err != nil {
			return []endpoints.Event{}, err
		}
		eventsList = append(eventsList, event)
	}

	return eventsList, nil
}

func (r *SQLiteRepo) DeleteEndpoint(id string) error {
	statemnent, err := r.db.Prepare("DELETE FROM endpoints WHERE id = ?")
	if err != nil {
		return err
	}
	defer statemnent.Close()

	_, err = statemnent.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLiteRepo) UpdateEndpoint(endpoint endpoints.Endpoint) error {
	statement, err := r.db.Prepare("UPDATE endpoints SET name = ?, url = ?, interval = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(endpoint.Name, endpoint.URL, endpoint.Interval, endpoint.ID)
	if err != nil {
		return err
	}

	return nil
}
