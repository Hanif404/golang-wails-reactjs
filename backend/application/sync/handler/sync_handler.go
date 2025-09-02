package synchandler

import (
	"golang-wails-reactjs/backend/domain/sync"
	"golang-wails-reactjs/backend/pkg/wrapper"
)

type Handler struct {
	service sync.Service
}

func New(s sync.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) PostData(email string) string {
	payload := sync.Sync{
		CreatedBy: email,
	}
	err := h.service.SaveData(&payload)
	if err != nil {
		return wrapper.Response(false, err.Error(), nil)
	}
	return wrapper.Response(true, "Save Data Success", nil)
}

func (h *Handler) GetDatas() string {
	result, err := h.service.GetDatas()
	if err != nil {
		return wrapper.Response(false, err.Error(), nil)
	}
	return wrapper.Response(true, "Get Data Success", result)
}

func (h *Handler) GetData(id int) string {
	result, err := h.service.GetData(id)
	if err != nil {
		return wrapper.Response(false, err.Error(), nil)
	}
	return wrapper.Response(true, "Get Data by Id Success", result)
}
