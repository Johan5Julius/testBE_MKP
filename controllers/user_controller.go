package controllers

import (
	"encoding/json"
	"net/http"
	"testMKP/helper"
	"testMKP/models"
	"testMKP/services"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusMethodNotAllowed,
			Status: "METHOD NOT ALLOWED",
			Error:  "Only POST method is allowed",
		})
		return
	}

	var request models.UserCreateRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "invalid request body: " + err.Error(),
		})
		return
	}

	user, err := u.UserService.Create(request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	})
}

func (u *UserController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusMethodNotAllowed,
			Status: "METHOD NOT ALLOWED",
			Error:  "Only POST method is allowed",
		})
		return
	}

	var request models.UserUpdateRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "invalid request body: " + err.Error(),
		})
		return
	}

	if request.ID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "ID tidak boleh kosong",
		})
		return
	}

	user, err := u.UserService.Update(request)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	})
}

func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusMethodNotAllowed,
			Status: "METHOD NOT ALLOWED",
			Error:  "Only POST method is allowed",
		})
		return
	}

	var request struct {
		ID int `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "invalid request body: " + err.Error(),
		})
		return
	}

	if request.ID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "ID tidak boleh kosong",
		})
		return
	}

	err := u.UserService.Delete(request.ID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Error:  err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "User berhasil dihapus",
	})
}

func (u *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusMethodNotAllowed,
			Status: "METHOD NOT ALLOWED",
			Error:  "Only POST method is allowed",
		})
		return
	}

	var request struct {
		ID int `json:"id"`
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "invalid request body: " + err.Error(),
		})
		return
	}

	if request.ID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "ID tidak boleh kosong",
		})
		return
	}

	user, err := u.UserService.FindById(request.ID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusNotFound,
			Status: "Not found",
			Error:  "User tidak di temukan",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	})
}

func (u *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusMethodNotAllowed,
			Status: "METHOD NOT ALLOWED",
			Error:  "Only POST method is allowed",
		})
		return
	}

	users, err := u.UserService.FindAll()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Error:  err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helper.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   users,
	})
	return
}
