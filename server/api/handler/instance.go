package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mi-bear/infra-control/api/adapter"
	"github.com/mi-bear/infra-control/api/presenter"
	"github.com/mi-bear/infra-control/interface/repository"
	"github.com/mi-bear/infra-control/usecase/instance"
)

func NewCreateInstanceHandler(
	db *sql.DB,
	hostRepo repository.TransactionalHostRepository,
	instanceRepo repository.TransactionalInstanceRepository,
	addressRepo repository.TransactionalAddressRepository,
	keyRepo repository.TransactionalKeyRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start transaction
		tx, err := db.BeginTx(r.Context(), &sql.TxOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		u := instance.NewCreateInstanceUseCase(
			hostRepo.WithTx(tx),
			instanceRepo.WithTx(tx),
			addressRepo.WithTx(tx),
			keyRepo.WithTx(tx),
		)

		input, err := adapter.NewCreateInstanceInputPortFromRequest(r)
		if err != nil {
			log.Println(err)
			if err := tx.Rollback(); err != nil {
				log.Println(err)
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		output, err := u.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			if err := tx.Rollback(); err != nil {
				log.Println(err)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tx.Commit(); err != nil {
			log.Println(err)
			if err := tx.Rollback(); err != nil {
				log.Println(err)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewCreateInstancePresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewGetInstanceHandler(u *instance.GetInstanceUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewGetInstanceInputPortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		output, err := u.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewGetInstancePresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewUpdateInstanceHandler(u *instance.UpdateInstanceUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewUpdateInstanceInputPortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		output, err := u.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewUpdateInstancePresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewDeleteInstanceHandler(
	db *sql.DB,
	instanceRepo repository.TransactionalInstanceRepository,
	keyRepo repository.TransactionalKeyRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start transaction
		tx, err := db.BeginTx(r.Context(), &sql.TxOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		u := instance.NewDeleteInstanceUseCase(
			instanceRepo.WithTx(tx),
			keyRepo.WithTx(tx),
		)

		input, err := adapter.NewDeleteInstanceInputPortFromRequest(r)
		if err != nil {
			log.Println(err)
			if err := tx.Rollback(); err != nil {
				log.Println(err)
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		output, err := u.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			if err := tx.Rollback(); err != nil {
				log.Println(err)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := tx.Commit(); err != nil {
			log.Println(err)
			if err := tx.Rollback(); err != nil {
				log.Println(err)
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewDeleteInstancePresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

func NewListInstancesHandler(u *instance.ListInstancesUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewListInstancesInputPortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		output, err := u.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := json.NewEncoder(w).Encode(presenter.NewListInstancesPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
