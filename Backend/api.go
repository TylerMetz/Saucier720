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

	router.HandleFunc("/Signup", makeHTTPHandleFunc(s.handleSignup))
	router.HandleFunc("/Login", makeHTTPHandleFunc(s.handleLogin))
	router.HandleFunc("/Pantry", makeHTTPHandleFunc(s.handleGetPantry))
	router.HandleFunc("/Recipes", makeHTTPHandleFunc(s.handleGetRecipes))
	

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleSignup(w http.ResponseWriter, r *http.Request) error {
	req := new(SignupRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		return err
	}

	account, err := NewAccount(req.UserName, req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil{
		return err
	}
	if err := s.store.PostSignup(account); err != nil {
		return err
	}

	resp := SignupResponse{
		Response: "User Successfully Created!",
	}

	return WriteJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	req := new(LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil{
		return err
	}

	verify := s.store.CheckPassword(req.UserName, req.Password)
	if(verify){
		resp := LoginResponse{
			// WE ACTUALLY NEED TO GENERATE A COOKIE
			Cookie: "GeneratedCookie",
		}
		return WriteJSON(w, http.StatusOK, resp)
	}

	return WriteJSON(w, http.StatusBadRequest, 0)
}

func (s *APIServer) handleGetPantry(w http.ResponseWriter, r *http.Request) error {
	pantry, err := s.store.GetPantry()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, pantry)
}

func (s *APIServer) handleGetRecipes(w http.ResponseWriter, r *http.Request) error{
	
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
