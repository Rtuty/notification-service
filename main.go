package main

import (
	"context"
	"notification-service/internal/api"
	"notification-service/internal/storage"
	"notification-service/pkg/logger"
)

func main() {
	// Создаем контекст и логгер
	ctx := context.Background()
	log := logger.GetLogger()
	log.Info("starting")

	// Получаем конфиг через viper и подключаемся к базе данных
	cfg, err := storage.GetStorageConfig()
	if err != nil {
		panic(err)
	}

	client, err := storage.NewClient(ctx, 5, cfg, log)
	if err != nil {
		panic(err)
	}

	api.Handle(ctx, storage.NewStorage(client), log)
}
