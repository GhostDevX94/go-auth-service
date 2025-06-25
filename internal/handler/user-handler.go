package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"user-service/internal/dto"
	"user-service/pkg"
)

func (h *Handler) UserRegister(w http.ResponseWriter, r *http.Request) {
	logrus.Info("🔄 Processing user registration request")

	var data dto.RegisterUser

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		logrus.WithError(err).Error("❌ Failed to decode registration request body")
		pkg.BadRequest(err, w)
		return
	}

	if validateError := h.Validator.Struct(&data); validateError != nil {
		logrus.WithError(validateError).Warn("⚠️ Registration validation failed")
		if processErr := pkg.StatusUnprocessableEntity(validateError, w); processErr != nil {
			pkg.BadRequest(processErr, w)
		}
		return
	}

	logrus.WithField("email", data.Email).Info("👤 Attempting to register new user")
	register, err := h.Services.UserService.Register(r.Context(), data)

	if err != nil {
		logrus.WithError(err).WithField("email", data.Email).Error("❌ User registration failed")
		pkg.BadRequest(err, w)
		return
	}

	logrus.WithField("user_id", register.Email).Info("✅ User registered successfully")
	err = pkg.WriteResponse(w, register)
	if err != nil {
		logrus.WithError(err).Error("❌ Failed to write registration response")
		pkg.BadRequest(err, w)
		return
	}
}

func (h *Handler) UserLogin(w http.ResponseWriter, r *http.Request) {
	logrus.Info("🔐 Processing user login request")

	var data dto.LoginUser

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		logrus.WithError(err).Error("❌ Failed to decode login request body")
		pkg.BadRequest(err, w)
		return
	}

	if validateError := h.Validator.Struct(&data); validateError != nil {
		logrus.WithError(validateError).Warn("⚠️ Login validation failed")
		if processErr := pkg.StatusUnprocessableEntity(validateError, w); processErr != nil {
			pkg.BadRequest(processErr, w)
		}
		return
	}

	logrus.WithField("email", data.Email).Info("🔑 Attempting user login")
	token, err := h.Services.UserService.Login(r.Context(), data)

	if err != nil {
		logrus.WithError(err).WithField("email", data.Email).Error("❌ User login failed")
		pkg.BadRequest(err, w)
		return
	}

	logrus.WithField("email", data.Email).Info("✅ User logged in successfully")
	response := dto.ResponseToken{Token: token}

	err = pkg.WriteResponse(w, response)
	if err != nil {
		logrus.WithError(err).Error("❌ Failed to write login response")
		pkg.BadRequest(err, w)
		return
	}
}
