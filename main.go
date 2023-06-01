package main

import (
	"context"
	"notification-service/internal/api"
	"notification-service/internal/storage"
)

func main() {
	ctx := context.Background()
	// Подключаемся к базе данных
	cfg, err := storage.GetStorageConfig()
	if err != nil {
		panic(err)
	}

	client, err := storage.NewClient(ctx, 5, cfg)
	if err != nil {
		panic(err)
	}

	api.Handle(context.TODO(), storage.NewStorage(client))
}
