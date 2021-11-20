package api

import (
	"database/sql"

	"github.com/gorilla/mux"

	"github.com/mi-bear/infra-control/api/handler"
	"github.com/mi-bear/infra-control/interface/repository"
	"github.com/mi-bear/infra-control/usecase/address"
	"github.com/mi-bear/infra-control/usecase/host"
	"github.com/mi-bear/infra-control/usecase/instance"
)

func BuildRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	buildProjectRoutes(r, db)
	return r
}

func buildProjectRoutes(r *mux.Router, db *sql.DB) {
	hostRepo := repository.NewHostRepository(db)
	instanceRepo := repository.NewInstanceRepository(db)
	addressRepo := repository.NewAddressRepository(db)
	keyRepo := repository.NewKeyRepository(db)

	r.Handle("/", handler.NewHelloHandler()).Methods("GET")

	// Host
	r.Handle("/hosts", handler.NewCreateHostHandler(
		host.NewCreateHostUseCase(hostRepo))).Methods("POST")

	// Address
	r.Handle("/addresses", handler.NewCreateAddressHandler(
		address.NewCreateAddressUseCase(addressRepo))).Methods("POST")

	// Instance
	r.Handle("/instances", handler.NewCreateInstanceHandler(
		db, instanceRepo, addressRepo, keyRepo)).Methods("POST")
	r.Handle("/instances/{id}", handler.NewGetInstanceHandler(
		instance.NewGetInstanceUseCase(instanceRepo))).Methods("GET")
	r.Handle("/instances/{id}/status/{state}", handler.NewUpdateInstanceHandler(
		instance.NewUpdateInstanceUseCase(instanceRepo))).Methods("PATCH")
	r.Handle("/instances/{id}", handler.NewDeleteInstanceHandler(
		db, instanceRepo, keyRepo)).Methods("DELETE")
	r.Handle("/instances", handler.NewListInstancesHandler(
		instance.NewListInstancesUseCase(instanceRepo))).Methods("GET").Queries("state", "{state}")
}
