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

var ClientData *entities.Client
var MailingData *entities.Mailing

func (db *dbConnect) GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := db.con.GetAllClients(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}

func (db *dbConnect) GetMailings(w http.ResponseWriter, r *http.Request) {
	mailings, err := db.con.GetAllMailings(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mailings)
}

func (db *dbConnect) CreateObject(w http.ResponseWriter, r *http.Request) {
	tbl := mux.Vars(r)["tbl"]

	switch tbl {
	case "client":
		if err := json.NewDecoder(r.Body).Decode(&ClientData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		if err := db.con.AddClient(r.Context(), ClientData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	case "mailing":
		if err := json.NewDecoder(r.Body).Decode(&MailingData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		if err := db.con.AddMailing(r.Context(), MailingData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (db *dbConnect) UpdateObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	tbl := vars["tbl"]

	switch tbl {
	case "clients":

		if err := json.NewDecoder(r.Body).Decode(&ClientData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		if err := db.con.UpdateClient(r.Context(), ClientData, id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	case "mailing":
		if err := json.NewDecoder(r.Body).Decode(&MailingData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		if err := db.con.UpdateMailing(r.Context(), MailingData, id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("строка таблицы %s с id: %s была обновлена", tbl, id))
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
