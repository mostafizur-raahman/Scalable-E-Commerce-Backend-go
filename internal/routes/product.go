package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce/internal/domain"
	"ecommerce/internal/services"
	"ecommerce/internal/utils"
)

type ProductHandler struct {
	svc *services.Service
}

func NewProductHandler(svc *services.Service) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	products, err := h.svc.GetAllProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.WriteJSON(w, http.StatusOK, products)
}

func (h *ProductHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, map[string]string{"error": "invalid product ID"})
		return
	}

	product, err := h.svc.GetProduct(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	utils.WriteJSON(w, http.StatusOK, product)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var product domain.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		utils.WriteError(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	if err := h.svc.CreateProduct(&product); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.WriteJSON(w, http.StatusCreated, product)
}
