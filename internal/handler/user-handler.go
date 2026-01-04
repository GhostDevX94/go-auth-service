package handler

import (
	"net/http"
	"user-service/internal/dto"
	"user-service/pkg"

	"github.com/sirupsen/logrus"
)

// UserRegister godoc
// @Summary Register a new user
// @Description Creates a new user account with provided credentials
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body dto.RegisterUser true "User registration data"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]interface{} "Bad request - validation error or user already exists"
// @Router /register [post]
func (h *Handler) UserRegister(w http.ResponseWriter, r *http.Request) {
	var data dto.RegisterUser

	if !pkg.DecodeAndValidate(&data, r, w) {
		return
	}

	register, err := h.Services.UserService.Register(r.Context(), data)

	if err != nil {
		logrus.WithError(err).WithField("email", data.Email).Error("User registration failed")
		pkg.BadRequest(err, w)
		return
	}

	err = pkg.WriteResponse(w, register)
	if err != nil {
		logrus.WithError(err).Error("Failed to write registration response")
		pkg.BadRequest(err, w)
		return
	}
}

// UserLogin godoc
// @Summary User login
// @Description Authenticates user and returns access and refresh tokens
// @Tags Authentication
// @Accept json
// @Produce json
// @Param credentials body dto.LoginUser true "Login credentials"
// @Success 200 {object} dto.ResponseToken
// @Failure 400 {object} map[string]interface{} "Bad request - invalid credentials"
// @Router /login [post]
func (h *Handler) UserLogin(w http.ResponseWriter, r *http.Request) {

	var data dto.LoginUser

	if !pkg.DecodeAndValidate(&data, r, w) {
		return
	}

	tokenPair, err := h.Services.UserService.Login(r.Context(), data)

	if err != nil {
		logrus.WithError(err).WithField("email", data.Email).Error("User login failed")
		pkg.BadRequest(err, w)
		return
	}

	err = pkg.WriteResponse(w, tokenPair)
	if err != nil {
		logrus.WithError(err).Error("Failed to write login response")
		pkg.BadRequest(err, w)
		return
	}
}

// UpdateToken godoc
// @Summary Refresh access token
// @Description Generates new access and refresh tokens using a valid refresh token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param token body dto.RefreshTokenRequest true "Refresh token"
// @Success 200 {object} dto.ResponseToken
// @Failure 400 {object} map[string]interface{} "Bad request - invalid or expired refresh token"
// @Router /update-token [post]
func (h *Handler) UpdateToken(w http.ResponseWriter, r *http.Request) {
	var data dto.RefreshTokenRequest

	if !pkg.DecodeAndValidate(&data, r, w) {
		return
	}

	tokenPair, err := h.Services.UserService.UpdateToken(r.Context(), data.RefreshToken)

	if err != nil {
		logrus.WithError(err).Error("Token refresh failed")
		pkg.BadRequest(err, w)
		return
	}

	err = pkg.WriteResponse(w, tokenPair)
	if err != nil {
		logrus.WithError(err).Error("Failed to write token refresh response")
		pkg.BadRequest(err, w)
		return
	}
}
