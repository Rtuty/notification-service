package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notification-service/internal/entities"
	"notification-service/internal/storage"

	"github.com/gorilla/mux"
)

type dbConnect struct {
	con storage.Storage
}

func (db *dbConnect) GetClients(w http.ResponseWriter, r *http.Request) {
	allClients, err := db.con.FindAllClients(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allClients)
}

func (db *dbConnect) DeleteObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	tbl := vars["tbl"]

	if err := db.con.DeleteRow(r.Context(), id, tbl); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("строка с id: %s была удалена из таблицы: %s", id, tbl))
}

func (db *dbConnect) UpdateClient(w http.ResponseWriter, r *http.Request) {
	var data *entities.Client
	id := mux.Vars(r)["id"]

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if err := db.con.UpdateClient(r.Context(), data, id); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("строка таблицы clients с id: %s была обновлена", id))
}
