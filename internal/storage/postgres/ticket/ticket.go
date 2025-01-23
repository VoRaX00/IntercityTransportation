package ticket

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"kursachDB/internal/domain/models"
)

type Ticket struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Ticket {
	return &Ticket{
		db: db,
	}
}

func (r *Ticket) Add(ticket models.Ticket) error {
	const op = "postgres.ticket.Add"
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	var id int
	query := `INSERT INTO tickets (cost, phone_number) VALUES ($1, $2) RETURNING id`
	err = tx.QueryRow(query, ticket.Cost, ticket.User.PhoneNumber).Scan(&id)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	for _, v := range ticket.Flight {
		res, err := checkPlace(tx, v.Id)
		if !res {
			_ = tx.Rollback()
			if err != nil {
				return fmt.Errorf("%s: %w", op, err)
			}
			return fmt.Errorf("%s: %w", op, "there are no seats left")
		}
	}

	query = `INSERT INTO tickets_flights (ticket_id, flight_id) VALUES ($1, $2)`
	for _, v := range ticket.Flight {
		_, err = tx.Exec(query, id, v.Id)
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return tx.Commit()
}

func checkPlace(tx *sqlx.Tx, flightId int64) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM tickets_flights WHERE flight_id = $1`
	err := tx.QueryRow(query, flightId).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("%s: %w", query, err)
	}

	var stateNumber string
	query = `SELECT state_number FROM flights WHERE id = $1`
	err = tx.QueryRow(query, flightId).Scan(&stateNumber)
	if err != nil {
		return false, fmt.Errorf("%s: %w", query, err)
	}

	var places int
	query = `SELECT models.count_places FROM buses 
    JOIN models ON models.model = buses.model
    WHERE state_number = $1`

	err = tx.QueryRow(query, stateNumber).Scan(&places)
	if err != nil {
		return false, fmt.Errorf("%s: %w", query, err)
	}
	return places > count, nil
}

func (r *Ticket) Delete(id int64) error {
	const op = "postgres.ticket.Delete"
	query := `DELETE FROM tickets WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (r *Ticket) GetAll() ([]models.Ticket, error) {
	const op = "storage.ticket.GetAll"

	tx, err := r.db.Beginx()
	if tx == nil {
		return nil, fmt.Errorf("%s: %w", op, "tx is nil")
	}

	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	query := `
	SELECT 
    	tickets.id AS "id", 
    	tickets.cost AS "cost", 
		tickets.phone_number AS "users.phone_number",
		users.fio AS "users.fio"
	FROM tickets
	JOIN users ON users.phone_number = tickets.phone_number`

	var tickets []models.Ticket
	err = r.db.Select(&tickets, query)
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i, v := range tickets {
		var flights []models.Flight
		query = `SELECT f.flight_id AS "id" FROM tickets_flights f WHERE ticket_id=$1`
		err = tx.Select(&flights, query, v.Id)
		if err != nil {
			_ = tx.Rollback()
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tickets[i].Flight = flights
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tickets, nil
}

func (r *Ticket) GetByUser(phoneNumber int64) ([]models.Ticket, error) {
	const op = "storage.ticket.GetByUser"

	tx, err := r.db.Beginx()
	if err != nil {
		if tx == nil {
			return nil, fmt.Errorf("%s: %w", op, "tx is nil")
		}
		_ = tx.Rollback()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	query := `
	SELECT 
    	tickets.id AS "id", 
    	tickets.cost AS "cost", 
		tickets.phone_number AS "users.phone_number",
		users.fio AS "users.fio"
	FROM tickets
	JOIN users ON users.phone_number = tickets.phone_number
	WHERE tickets.phone_number = $1`

	var tickets []models.Ticket
	err = r.db.Select(&tickets, query, phoneNumber)
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	for i, v := range tickets {
		var flights []models.Flight
		query = `SELECT f.flight_id AS "id" FROM tickets_flights f WHERE ticket_id=$1`
		err = tx.Select(&flights, query, v.Id)
		if err != nil {
			_ = tx.Rollback()
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		tickets[i].Flight = flights
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tickets, nil
}
