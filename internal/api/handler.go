package api

import (
	"context"
	"fmt"
	"net/http"
	"notification-service/internal/storage"

	"github.com/gorilla/mux"
)

func Handle(ctx context.Context, conenct storage.Storage) {
	router := mux.NewRouter()
	con := &dbConnect{con: conenct}

	// Определение маршрутов и обработчиков запросов
	router.HandleFunc("/client", con.GetClients).Methods("GET")
	router.HandleFunc("/client/{id}", con.UpdateClient).Methods("UPDATE")
	router.HandleFunc("/delete/{tbl}/{id}", con.DeleteObject).Methods("DELETE")

	// Запуск сервера
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
