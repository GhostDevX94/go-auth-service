package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"user-service/internal/request"
)

func (h *Handler) UserRegister(w http.ResponseWriter, r *http.Request) {

	var data request.RegisterUser

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		request.BadRequest(err, w)
		return
	}
	log.Println(data)
	if validateError := h.Validator.Struct(&data); validateError != nil {
		if processErr := request.StatusUnprocessableEntity(validateError, w); processErr != nil {
			request.BadRequest(processErr, w)
		}
		return
	}

	fmt.Println(data)

}
