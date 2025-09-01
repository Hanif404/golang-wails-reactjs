package studenthandler

import (
	"golang-wails-reactjs/backend/domain/student"
	"golang-wails-reactjs/backend/pkg/validator"
	"golang-wails-reactjs/backend/pkg/wrapper"
)

type Handler struct {
	service student.Service
}

func New(s student.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) PostData(payload student.Student) string {
	//validation payload
	errValid := validator.ValidateStruct(payload)
	if errValid != nil {
		return wrapper.Response(false, errValid, nil)
	}

	err := h.service.SaveData(&payload)
	if err != nil {
		return wrapper.Response(false, err.Error(), nil)
	}
	return wrapper.Response(true, "Save Data Success", nil)
}

func (h *Handler) GetDatas() string {
	result, err := h.service.GetStudents()
	if err != nil {
		return wrapper.Response(false, err.Error(), nil)
	}
	return wrapper.Response(true, "Get Data Success", result)
}

func (h *Handler) GetData(id int) string {
	result, err := h.service.GetStudent(id)
	if err != nil {
		return wrapper.Response(false, err.Error(), nil)
	}
	return wrapper.Response(true, "Get Data by Id Success", result)
}

func (h *Handler) DeleteData(id int) string {
	err := h.service.DeleteData(id)
	if err != nil {
		return wrapper.Response(false, err.Error(), nil)
	}
	return wrapper.Response(true, "Delete Success", nil)
}
