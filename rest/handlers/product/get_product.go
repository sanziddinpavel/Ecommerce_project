package product

import (
	"Ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		util.SendError(w, http.StatusBadRequest, "invalid req body")
		return
	}
	product, err := h.productRepo.Get(pid)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if product == nil {
		util.SendError(w, http.StatusNotFound, "product not found")
		return

	}

	util.SendData(w, http.StatusOK, product)
}
