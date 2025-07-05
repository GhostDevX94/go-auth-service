package handler

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"user-service/internal/dto"
	"user-service/pkg"
)

func (h *Handler) UserRegister(w http.ResponseWriter, r *http.Request) {
	var data dto.RegisterUser

	if !pkg.DecodeAndValidate(&data, r, w) {
		return
	}

	register, err := h.Services.UserService.Register(r.Context(), data)

	if err != nil {
		logrus.WithError(err).WithField("email", data.Email).Error("❌ User registration failed")
		pkg.BadRequest(err, w)
		return
	}

	err = pkg.WriteResponse(w, register)
	if err != nil {
		logrus.WithError(err).Error("❌ Failed to write registration response")
		pkg.BadRequest(err, w)
		return
	}
}

func (h *Handler) UserLogin(w http.ResponseWriter, r *http.Request) {

	var data dto.LoginUser

	if !pkg.DecodeAndValidate(&data, r, w) {
		return
	}

	token, err := h.Services.UserService.Login(r.Context(), data)

	if err != nil {
		logrus.WithError(err).WithField("email", data.Email).Error("❌ User login failed")
		pkg.BadRequest(err, w)
		return
	}

	response := dto.ResponseToken{Token: token}

	err = pkg.WriteResponse(w, response)
	if err != nil {
		logrus.WithError(err).Error("❌ Failed to write login response")
		pkg.BadRequest(err, w)
		return
	}
}
