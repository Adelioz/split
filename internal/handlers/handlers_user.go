package handlers

import (
	"errors"
	"net/http"

	"github.com/Adelioz/split/internal/models"
	"github.com/Adelioz/split/internal/service"
	"github.com/Adelioz/split/pkg/logging"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func AddUser(s *service.Service) http.Handler {
	type addUserRequest struct {
		ID       string `json:"id" validate:"required"`
		Username string `json:"username" validate:"required"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var request addUserRequest
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

		user := models.User{
			ID:       request.ID,
			Username: request.Username,
		}
		err = s.AddUser(ctx, user)

		// TODO: - constant errors handling (repo.go)
		if err != nil {
			logging.S(ctx).Warnf("Failed to AddUser: %s.", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		WriteJSONResponse(ctx, w, http.StatusOK, user)
	})

}

func UpdateUser(s *service.Service) http.Handler {
	type updateUserRequest struct {
		ID       string `json:"id"`
		Username string `json:"username" validate:"required"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var request updateUserRequest
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

		user := models.User{
			ID:       request.ID,
			Username: request.Username,
		}
		err = s.UpdateUser(ctx, user)

		// TODO: - constant errors handling (repo.go)
		if err != nil {
			logging.S(ctx).Warnf("Failed to UpdateUser: %s.", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		WriteJSONResponse(ctx, w, http.StatusOK, user)
	})
}

func GetUser(s *service.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, ok := mux.Vars(r)["id"]

		if !ok {
			logging.S(ctx).Warn("There are no id found.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := s.GetUser(ctx, id)

		if errors.Is(err, service.ErrNotFound) {
			logging.S(ctx).Warn("User not found.")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err != nil {
			logging.S(ctx).Warnf("Couldn't get User: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		WriteJSONResponse(ctx, w, http.StatusOK, user)
	})
}
