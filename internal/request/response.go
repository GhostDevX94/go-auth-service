package request

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func StatusUnprocessableEntity(err error, w http.ResponseWriter) error {
	var invalidValidationError *validator.InvalidValidationError
	var responseErrors map[string]string

	w.Header().Set("Content-Type", "application/json")
	if errors.As(err, &invalidValidationError) {
		return err
	}

	var validateErrs validator.ValidationErrors
	if errors.As(err, &validateErrs) {
		responseErrors = make(map[string]string)
		for _, e := range validateErrs {
			responseErrors[e.Field()] = e.Error()
		}

		w.WriteHeader(http.StatusUnprocessableEntity)
		if writeErr := WriteResponse(w, responseErrors); writeErr != nil {
			return fmt.Errorf("failed to write response: %w", writeErr)
		}
		return nil

	}

	return fmt.Errorf("unknown error type: %w", err)

}

func BadRequest(err error, w http.ResponseWriter) {

	w.WriteHeader(http.StatusBadRequest)

	if writeErr := WriteResponse(w, err.Error()); writeErr != nil {
		_ = fmt.Errorf("failed to write response: %w", writeErr)
	}

}

func WriteResponse(w http.ResponseWriter, data any) error {
	marshal, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to write response: %w", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err = w.Write(marshal)
	if err != nil {
		return err
	}

	return nil
}
