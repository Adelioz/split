package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Adelioz/split/pkg/logging"
	"github.com/go-playground/validator/v10"
)

func WriteJSONResponse(ctx context.Context, w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		logging.S(ctx).Errorf("Failed to write response: %w", err)
	}
}

func ReadJSONRequest(ctx context.Context, r *http.Request, dst any) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return fmt.Errorf("json decoding: %w", err)
	}

	err = validator.New(validator.WithRequiredStructEnabled()).Struct(dst)
	if err != nil {
		return fmt.Errorf("validation: %w", err)
	}

	return nil
}
