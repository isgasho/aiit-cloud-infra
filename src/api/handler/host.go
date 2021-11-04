package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mi-bear/infra-control/api/adapter"
	"github.com/mi-bear/infra-control/api/presenter"
	"github.com/mi-bear/infra-control/usecase/host"
)

func NewCreateHostHandler(u *host.CreateHostUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		input, err := adapter.NewCreateHostInputPortFromRequest(r)
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
		if err := json.NewEncoder(w).Encode(presenter.NewCreateHostPresenter(output)); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
