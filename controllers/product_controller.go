package controllers

import (
	"encoding/json"
	"net/http"
	"testMKP/helper"
	"testMKP/models"
	"testMKP/services"
)

type ProductController struct {
	ProductService services.ProductService
}

func NewProductController(productService services.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (p *ProductController) Create(w http.ResponseWriter, r *http.Request) {
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

	var request models.ProductCreateRequest

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

	product, err := p.ProductService.Create(request)
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
		Data:   product,
	})
}

func (p *ProductController) Update(w http.ResponseWriter, r *http.Request) {
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

	var request models.ProductUpdateRequest

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

	if request.ID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "ID tidak boleh kosong",
		})
		return
	}

	product, err := p.ProductService.Update(request)
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
		Data:   product,
	})
}

func (p *ProductController) Delete(w http.ResponseWriter, r *http.Request) {
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
		ID string `json:"id"`
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

	if request.ID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "ID tidak boleh kosong",
		})
		return
	}

	err := p.ProductService.Delete(request.ID)
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
		Data:   "Product berhasil dihapus",
	})
}

func (p *ProductController) FindById(w http.ResponseWriter, r *http.Request) {
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
		ID string `json:"id"`
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

	if request.ID == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helper.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Error:  "ID tidak boleh kosong",
		})
		return
	}

	product, err := p.ProductService.FindById(request.ID)
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
		Data:   product,
	})
}

func (p *ProductController) FindAll(w http.ResponseWriter, r *http.Request) {
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

	product, err := p.ProductService.FindAll()
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
		Data:   product,
	})
	return
}
