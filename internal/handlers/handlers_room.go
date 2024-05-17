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

func AddRoom(s *service.Service) http.Handler {
	type addRoomRequest struct {
		ID   string `json:"id" validate:"required"`
		Name string `json:"name" validate:"required"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var request addRoomRequest
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

		room := models.Room{
			ID:   request.ID,
			Name: request.Name,
		}
		err = s.AddRoom(ctx, room)

		// TODO: - constant errors handling (repo.go)
		if err != nil {
			logging.S(ctx).Warnf("Failed to AddRoom: %s.", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		WriteJSONResponse(ctx, w, http.StatusOK, room)
	})

}

func UpdateRoom(s *service.Service) http.Handler {
	type updateRoomRequest struct {
		ID   string `json:"id" validate:"required"`
		Name string `json:"name" validate:"required"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var request updateRoomRequest
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

		room := models.Room{
			ID:   request.ID,
			Name: request.Name,
		}
		err = s.UpdateRoom(ctx, room)

		// TODO: - constant errors handling (repo.go)
		if err != nil {
			logging.S(ctx).Warnf("Failed to UpdateRoom: %s.", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		WriteJSONResponse(ctx, w, http.StatusOK, room)
	})
}

func GetRoom(s *service.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, ok := mux.Vars(r)["id"]

		if !ok {
			logging.S(ctx).Warn("There are no id found.")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		room, err := s.GetRoom(ctx, id)

		if err != nil {
			logging.S(ctx).Warn("Couldn't get Room.")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		WriteJSONResponse(ctx, w, http.StatusOK, room)
	})
}
