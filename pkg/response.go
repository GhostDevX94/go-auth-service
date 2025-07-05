package pkg

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

	response := map[string]string{
		"message": err.Error(),
	}

	if writeErr := WriteResponse(w, response); writeErr != nil {
		_ = fmt.Errorf("failed to write response: %w", writeErr)
	}

}

var v *validator.Validate

func init() {
	v = validator.New(validator.WithRequiredStructEnabled())
}

func DecodeAndValidate(data any, r *http.Request, w http.ResponseWriter) bool {

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		BadRequest(err, w)
		return false
	}

	if validateError := v.Struct(data); validateError != nil {
		err := StatusUnprocessableEntity(validateError, w)
		if err != nil {
			BadRequest(err, w)
			return false
		}
		return false
	}

	return true
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
