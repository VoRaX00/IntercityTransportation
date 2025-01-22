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
	query := `INSERT INTO tickets (cost, phone_number) VALUES ($1, $2)`
	_, err := r.db.Exec(query, ticket.Cost, ticket.User.PhoneNumber)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
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
	query := `
	SELECT 
    	id, 
    	cost, 
		tickets.phone_number,
		users.fio
	FROM tickets
	JOiN users ON users.phone_number = tickets.phone_number`

	var tickets []models.Ticket
	err := r.db.Select(&tickets, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tickets, nil
}

func (r *Ticket) GetByUser(phoneNumber int64) ([]models.Ticket, error) {
	const op = "storage.ticket.GetAll"
	query := `
	SELECT 
    	id, 
    	cost, 
		tickets.phone_number,
		users.fio
	FROM tickets
	JOiN users ON users.phone_number = tickets.phone_number
	WHERE tickets.phone_number = $1`

	var tickets []models.Ticket
	err := r.db.Select(&tickets, query, phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return tickets, nil
}
