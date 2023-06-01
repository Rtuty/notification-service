package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notification-service/internal/entities"
	"notification-service/internal/storage"
	"notification-service/pkg/logger"
	"sync"

	"github.com/gorilla/mux"
)

type server struct {
	con         storage.Storage
	log         *logger.Logger
	clientData  *entities.Client
	mailingData *entities.Mailing
	mu          sync.RWMutex
}

func NewServer(con storage.Storage, log *logger.Logger) *server {
	return &server{
		con: con,
		log: log,
	}
}

func (s *server) GetClients(w http.ResponseWriter, r *http.Request) {
	clients, err := s.con.GetAllClients(r.Context())
	if err != nil {
		s.log.Panic(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}

func (s *server) GetMailings(w http.ResponseWriter, r *http.Request) {
	mailings, err := s.con.GetAllMailings(r.Context())
	if err != nil {
		s.log.Panic(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mailings)
}

func (s *server) CreateObject(w http.ResponseWriter, r *http.Request) {
	tbl := mux.Vars(r)["tbl"]

	s.mu.Lock()
	defer s.mu.Unlock()

	switch tbl {
	case "client":
		if err := json.NewDecoder(r.Body).Decode(&s.clientData); err != nil {
			s.log.Panic(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		if err := s.con.AddClient(r.Context(), s.clientData); err != nil {
			s.log.Panic(err)
			w.WriteHeader(http.StatusBadRequest)
		}
	case "mailing":
		if err := json.NewDecoder(r.Body).Decode(&s.mailingData); err != nil {
			s.log.Panic(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		if err := s.con.AddMailing(r.Context(), s.mailingData); err != nil {
			s.log.Panic(err)
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (s *server) UpdateObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	tbl := vars["tbl"]

	s.mu.Lock()
	defer s.mu.Unlock()

	switch tbl {
	case "clients":

		if err := json.NewDecoder(r.Body).Decode(&s.clientData); err != nil {
			s.log.Panic(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		if err := s.con.UpdateClient(r.Context(), s.clientData, id); err != nil {
			s.log.Panic(err)
			w.WriteHeader(http.StatusBadRequest)
		}
	case "mailing":
		if err := json.NewDecoder(r.Body).Decode(&s.mailingData); err != nil {
			s.log.Panic(err)
			w.WriteHeader(http.StatusBadRequest)
		}

		if err := s.con.UpdateMailing(r.Context(), s.mailingData, id); err != nil {
			s.log.Panic(err)
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("строка таблицы %s с id: %s была обновлена", tbl, id))
}

func (s *server) DeleteObject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	tbl := vars["tbl"]

	if err := s.con.DeleteRow(r.Context(), id, tbl); err != nil {
		s.log.Panic(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(fmt.Sprintf("строка с id: %s была удалена из таблицы: %s", id, tbl))
}
