package storage

import (
	"context"
	"errors"
	"fmt"
	"notification-service/internal/entities"

	"github.com/jackc/pgconn"
)

type Storage interface {
	AddClient(ctx context.Context, cl *entities.Client) error
	UpdateClient(ctx context.Context, id string, cl *entities.Client) error // Todo update client
	FindAllClients(ctx context.Context) (cl []entities.Client, err error)

	AddMailing(ctx context.Context, m *entities.Mailing) error
	UpdateMailing(ctx context.Context, id string, m *entities.Mailing) error // Todo update mailing
	GetMailingStatistics(ctx context.Context) (res string, err error)        //todo result parameters

	DeleteRow(ctx context.Context, id string, tbl string) error
}
type db struct {
	client Client
}

func NewStorage(client Client) Storage {
	return &db{client: client}
}

var PgErr *pgconn.PgError
var ErrQ error

// Функции для объекта клиента
func (db *db) AddClient(ctx context.Context, cl *entities.Client) error {
	q := `
		insert into clients (phone_number, phone_code, tag, time_zone)
		values ($1, $2, $3, $4) returning id
		`
	ErrQ = db.client.QueryRow(ctx, q, cl.PhoneCode, cl.PhoneCode, cl.Tag, cl.TimeZone).Scan(&cl.ID)

	if errors.Is(ErrQ, PgErr) {
		PgErr = ErrQ.(*pgconn.PgError)
		newErr := fmt.Errorf(fmt.Sprintf("sql error: %s,  Detail: %s, Where: %s, Code: %s, SQLState: %s", PgErr.Message, PgErr.Detail, PgErr.Where, PgErr.Code, PgErr.SQLState()))
		return newErr
	}
	return nil
}

func (db *db) UpdateClient(ctx context.Context, id string, cl *entities.Client) error { return nil }

func (db *db) FindAllClients(ctx context.Context) ([]entities.Client, error) {
	q := `select id, phone_number, phone_code, tag, time_zone from clients`

	rows, err := db.client.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("sql select all clients error: %v", err))
	}

	clients := make([]entities.Client, 0)

	for rows.Next() {
		var cln entities.Client

		err = rows.Scan(&cln.ID, &cln.PhoneNumber, &cln.PhoneCode, &cln.Tag, &cln.TimeZone)
		if err != nil {
			return nil, err
		}

		clients = append(clients, cln)
	}

	return clients, nil
}

// Функции для объекта рассылка
func (db *db) AddMailing(ctx context.Context, m *entities.Mailing) error {
	q := `
		insert into clients (start_time, finish_time, message, filter)
		values ($1, $2, $3, $4) returning id
	`
	ErrQ = db.client.QueryRow(ctx, q, m.StartTime, m.FinishTime, m.Message, m.Filter).Scan(&m.ID)

	if errors.Is(ErrQ, PgErr) {
		PgErr = ErrQ.(*pgconn.PgError)
		newErr := fmt.Errorf(fmt.Sprintf("sql error: %s,  Detail: %s, Where: %s, Code: %s, SQLState: %s", PgErr.Message, PgErr.Detail, PgErr.Where, PgErr.Code, PgErr.SQLState()))
		return newErr
	}
	return nil
}

func (db *db) UpdateMailing(ctx context.Context, id string, m *entities.Mailing) error { return nil }

func (db *db) GetAllMailings(ctx context.Context) ([]entities.Mailing, error) {
	q := `select id, start_time, finish_time, message, filter from mailings`

	rows, err := db.client.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("sql select all mailings error: %v", err))
	}

	mailings := make([]entities.Mailing, 0)

	for rows.Next() {
		var m entities.Mailing

		err = rows.Scan(&m.ID, &m.StartTime, &m.FinishTime, &m.Message, &m.Filter)
		if err != nil {
			return nil, err
		}

		mailings = append(mailings, m)
	}

	return mailings, nil
}

func (db *db) GetMailingStatistics(ctx context.Context) (string, error) { return "", nil }

// Общесистемные функции
func (db *db) DeleteRow(ctx context.Context, id string, tbl string) error {
	var q string
	switch tbl {
	case "clients":
		q = `delete from clients where id = $1`
	case "mailings":
		q = `delete from mailings where id = $1`
	}
	if err := db.client.QueryRow(ctx, q); err != nil {
		return fmt.Errorf(fmt.Sprintf("sql delete %s error: %v", tbl, err))
	}
	return nil
}
