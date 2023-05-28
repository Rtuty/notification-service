package storage

import (
	"context"
	"notification-service/internal/entities"
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
