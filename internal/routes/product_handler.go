package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ecommerce/internal/domain/product"
	"ecommerce/internal/services"
	"ecommerce/internal/utils"
)

type ProductHandler struct {
	svc *services.ProductService
}

func NewProductHandler(svc *services.ProductService) *ProductHandler {
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
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	p, err := h.svc.GetProduct(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	utils.WriteJSON(w, http.StatusOK, p)
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var p product.Product
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		utils.WriteError(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return
	}
	if err := h.svc.CreateProduct(&p); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	utils.WriteJSON(w, http.StatusCreated, p)
}
