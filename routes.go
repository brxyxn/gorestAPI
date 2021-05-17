package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

const (
	host     = "queenie.db.elephantsql.com"
	port     = "5432"
	user     = "xisreddd"
	password = "c3aGVFPauG2BZZUzM5TiuI42dFYW8tSZ"
	dbname   = "xisreddd"
)

func (a *App) Initialize(user, password, dbname string) {
	// connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	// usr := "xyxdev"
	// pwd := "admin"
	// dbn := "go-rest-api"
	connectionStr := fmt.Sprintf("host=%s port=%v user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("pgx", connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":3005", a.Router))
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/v1/products/new", a.newRestaurant).Methods("POST")
	a.Router.HandleFunc("/api/v1/products", a.getRestaurants).Methods("GET")
	a.Router.HandleFunc("/api/v1/product/{id:[0-9]+}", a.getRestaurant).Methods("GET")
	a.Router.HandleFunc("/api/v1/product/{id:[0-9]+}", a.updateRestaurant).Methods("PUT")
	a.Router.HandleFunc("/api/v1/product/{id:[0-9]+}", a.deleteRestaurant).Methods("DELETE")

	a.Router.HandleFunc("/api/v1/product/{restaurant_id:[0-9]+}/comments/new", a.newComment).Methods("GET")
	a.Router.HandleFunc("/api/v1/product/{restaurant_id:[0-9]+}/comments", a.getComments).Methods("GET")
	a.Router.HandleFunc("/api/v1/product/{restaurant_id:[0-9]+}/comments/{id:[0-9]+}", a.updateComment).Methods("GET")
	a.Router.HandleFunc("/api/v1/product/{restaurant_id:[0-9]+}/comments/{id:[0-9]+}", a.deleteComment).Methods("GET")
}
