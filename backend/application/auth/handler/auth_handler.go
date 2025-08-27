package authhandler

import (
	"golang-wails-reactjs/backend/domain/auth"
	"golang-wails-reactjs/backend/pkg/validator"
	"golang-wails-reactjs/backend/pkg/wrapper"
)

type Handler struct {
	service auth.Service
}

func New(s auth.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) PostLogin(payload auth.Login) string {
	//validation payload
	errValid := validator.ValidateStruct(payload)
	if errValid != nil {
		return wrapper.Response(false, "email atau password tidak sesuai", nil)
	}

	err := h.service.Login(&payload)
	if err != nil {
		return wrapper.Response(false, err.Error(), nil)
	}
	return wrapper.Response(true, "Login success", nil)
}
