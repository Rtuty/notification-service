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
	con := NewServer(conenct)

	// Определение маршрутов и обработчиков запросов
	router.HandleFunc("/clients", con.GetClients).Methods("GET")
	router.HandleFunc("/mailings", con.GetMailings).Methods("GET")
	router.HandleFunc("/add/{tbl}", con.CreateObject).Methods("POST")

	router.HandleFunc("/update/{tbl}/{id}", con.UpdateObject).Methods("PUT")
	router.HandleFunc("/delete/{tbl}/{id}", con.DeleteObject).Methods("DELETE")

	// Запуск сервера
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}
