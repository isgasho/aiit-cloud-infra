package api

import (
	"database/sql"

	"github.com/gorilla/mux"
)

func BuildRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	buildProjectRoutes(r, db)
	return r
}

func buildProjectRoutes(r *mux.Router, db *sql.DB) {

}
