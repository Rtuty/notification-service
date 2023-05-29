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
	UpdateClient(ctx context.Context, cl *entities.Client) error
	DeleteClient(ctx context.Context, cl *entities.Client) error
	FindAllClients(ctx context.Context) (cl []entities.Client, err error)

	AddMailing(ctx context.Context, m *entities.Mailing) error
	UpdateMailing(ctx context.Context, m *entities.Mailing) error
	DeleteMailing(ctx context.Context, m *entities.Mailing) error
	GetMailingStatistics(ctx context.Context) (res string, err error) //todo result parameters
}
type db struct {
	client Client
}

func NewStorage(client Client) Storage {
	return &db{client: client}
}

func (db *db) AddClient(ctx context.Context, cl *entities.Client) error {
	var pgErr *pgconn.PgError
	var errQ error

	q := `
		insert into cliens (phone_number, phone_code, tag, time_zone)
		values ($1, $2, $3, $4) returning id
		`
	errQ = db.client.QueryRow(ctx, q, cl.PhoneCode, cl.PhoneCode, cl.Tag, cl.TimeZone).Scan(&cl.ID)

	if errors.Is(errQ, pgErr) {
		pgErr = errQ.(*pgconn.PgError)
		newErr := fmt.Errorf(fmt.Sprintf("sql error: %s,  Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
		return newErr
	}
	return nil
}

func (db *db) UpdateClient(ctx context.Context, cl *entities.Client) error {
	return nil
}

func (db *db) DeleteClient(ctx context.Context, cl *entities.Client) error { return nil }
func (db *db) FindAllClients(ctx context.Context) ([]entities.Client, error) {
	return []entities.Client{}, nil
}

func (db *db) AddMailing(ctx context.Context, m *entities.Mailing) error    { return nil }
func (db *db) UpdateMailing(ctx context.Context, m *entities.Mailing) error { return nil }
func (db *db) DeleteMailing(ctx context.Context, m *entities.Mailing) error { return nil }
func (db *db) GetMailingStatistics(ctx context.Context) (string, error)     { return "", nil } //todo result parameters
