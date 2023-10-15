package main

import (
	"encoding/json"
	_"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store Storage
}

func NewMealDealzServer(listenAddr string, store Storage) *APIServer{
	return &APIServer{
		listenAddr: listenAddr,
		store: store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/Pantry", makeHTTPHandleFunc(s.handleGetPantry))
	router.HandleFunc("/Signup", makeHTTPHandleFunc(s.handleSignup))

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleGetPantry(w http.ResponseWriter, r *http.Request) error {
	pantry, err := s.store.GetPantry()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, pantry)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
