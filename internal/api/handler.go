package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"notification-service/internal/storage"

	"github.com/gorilla/mux"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cfg, err := storage.GetStorageConfig()
	if err != nil {
		w.WriteHeader(400)
		panic(err)
	}

	client, err := storage.NewClient(ctx, 5, cfg)
	if err != nil {
		w.WriteHeader(400)
		panic(err)
	}

	allClients, err := storage.NewStorage(client).FindAllClients(ctx)
	if err != nil {
		w.WriteHeader(400)
		panic(err)
	}
	json.NewEncoder(w).Encode(allClients)
}

func Handle(ctx context.Context) {
	router := mux.NewRouter()

	// Определение маршрутов и обработчиков запросов
	router.HandleFunc("/client", GetClients).Methods("GET")

	// Запуск сервера
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
