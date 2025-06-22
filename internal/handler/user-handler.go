package handler

import (
	"encoding/json"
	"net/http"
	"user-service/internal/dto"
	"user-service/pkg"
)

func (h *Handler) UserRegister(w http.ResponseWriter, r *http.Request) {

	var data dto.RegisterUser

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		pkg.BadRequest(err, w)
		return
	}
	if validateError := h.Validator.Struct(&data); validateError != nil {
		if processErr := pkg.StatusUnprocessableEntity(validateError, w); processErr != nil {
			pkg.BadRequest(processErr, w)
		}
		return
	}

	register, err := h.Services.UserService.Register(r.Context(), data)

	if err != nil {
		pkg.BadRequest(err, w)
		return
	}

	err = pkg.WriteResponse(w, register)
	if err != nil {
		pkg.BadRequest(err, w)
		return
	}

}

func (h *Handler) UserLogin(w http.ResponseWriter, r *http.Request) {
	var data dto.LoginUser

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		pkg.BadRequest(err, w)
		return
	}
	if validateError := h.Validator.Struct(&data); validateError != nil {
		if processErr := pkg.StatusUnprocessableEntity(validateError, w); processErr != nil {
			pkg.BadRequest(processErr, w)
		}
		return
	}

	token, err := h.Services.UserService.Login(r.Context(), data)

	if err != nil {
		pkg.BadRequest(err, w)
		return
	}

	response := dto.ResponseToken{Token: token}

	err = pkg.WriteResponse(w, response)
	if err != nil {
		pkg.BadRequest(err, w)
		return
	}
}
