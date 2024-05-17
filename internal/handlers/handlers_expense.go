package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/Adelioz/split/internal/models"
	"github.com/Adelioz/split/internal/service"
	"github.com/Adelioz/split/pkg/logging"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func AddExpense(s *service.Service) http.Handler {

	type addExpenseRequest struct {
		ID        string    `json:"id" validate:"required"`
		RoomID    string    `json:"roomId"` // TBD required(?)
		DaddyID   string    `json:"daddyId" validate:"required"`
		Currency  string    `json:"currency" validate:"required"`
		Tag       string    `json:"tag"`
		CreatedAt time.Time `json:"createdAt"` // TBD required(?)
		Title     string    `json:"title" validate:"required"`
		Desc      string    `json:"desc"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var request addExpenseRequest
		err := ReadJSONRequest(ctx, r, &request)

		var validationErr validator.ValidationErrors

		// TODO: Error is EOF
		// if errors.Is(err, io.EOF) {
		// 	logging.S(ctx).Debugf("IO Error: %s.", err)
		// 	w.WriteHeader(http.StatusBadRequest)
		// }
		// hendle in Utils.go(?)

		if errors.As(err, &validationErr) {
			logging.S(ctx).Debugf("Validation error: %s.", validationErr)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err != nil {
			logging.S(ctx).Warnf("Failed to read request: %s.", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		exp := models.Expense{
			ID:        request.ID,
			RoomID:    request.RoomID,
			DaddyID:   request.DaddyID,
			Currency:  request.Currency,
			Tag:       request.Tag,
			CreatedAt: request.CreatedAt,
			Title:     request.Title,
			Desc:      request.Desc,
		}
		err = s.AddExpense(ctx, exp)

		// TODO: - constant errors handling (repo.go)
		if err != nil {
			logging.S(ctx).Warnf("Failed to AddExpense: %s.", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		WriteJSONResponse(ctx, w, http.StatusOK, exp)
	})

}

func UpdateExpense(s *service.Service) http.Handler {

	type updateExpenseRequest struct {
		ID        string    `json:"id" validate:"required"`
		RoomID    string    `json:"roomId"`
		DaddyID   string    `json:"daddyId"`
		Currency  string    `json:"currency"`
		Tag       string    `json:"tag"`
		CreatedAt time.Time `json:"createdAt"`
		Title     string    `json:"title"`
		Desc      string    `json:"desc"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var request updateExpenseRequest
		err := ReadJSONRequest(ctx, r, &request)

		var validationErr validator.ValidationErrors

		// TODO: Error is EOF
		// if errors.Is(err, io.EOF) {
		// 	logging.S(ctx).Debugf("IO Error: %s.", err)
		// 	w.WriteHeader(http.StatusBadRequest)
		// }
		// hendle in Utils.go(?)

		if errors.As(err, &validationErr) {
			logging.S(ctx).Debugf("Validation error: %s.", validationErr)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err != nil {
			logging.S(ctx).Warnf("Failed to read request: %s.", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		exp := models.Expense{
			ID:        request.ID,
			RoomID:    request.RoomID,
			DaddyID:   request.DaddyID,
			Currency:  request.Currency,
			Tag:       request.Tag,
			CreatedAt: request.CreatedAt,
			Title:     request.Title,
			Desc:      request.Desc,
		}
		err = s.UpdateExpense(ctx, exp)

		// TODO: - constant errors handling (repo.go)
		if err != nil {
			logging.S(ctx).Warnf("Failed to AddExpense: %s.", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		WriteJSONResponse(ctx, w, http.StatusOK, exp)
	})
}

func GetExpense(s *service.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, ok := mux.Vars(r)["id"]

		if !ok {
			logging.S(ctx).Warn("There are no id found.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		exp, err := s.GetExpense(ctx, id)

		if err != nil {
			logging.S(ctx).Warn("Couldn't get Expense.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		WriteJSONResponse(ctx, w, http.StatusOK, exp)
	})
}

func DeleteExpense(s *service.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, ok := mux.Vars(r)["id"]

		if !ok {
			logging.S(ctx).Warn("There are no id found.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		exp, err := s.DeleteExpense(ctx, id)

		if err != nil {
			logging.S(ctx).Warn("Couldn't delete Expense.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		WriteJSONResponse(ctx, w, http.StatusOK, exp)
	})
}
