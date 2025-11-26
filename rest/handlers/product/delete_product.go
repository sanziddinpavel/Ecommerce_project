package product

import (
	"Ecommerce/util"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		fmt.Println(err)

		util.SendError(w, http.StatusBadRequest, "invalid product id")
		return

	}
	err = h.productRepo.Delete(pid)
	if err != nil {

		util.SendError(w, http.StatusInternalServerError, "internal server error")
		return

	}
	util.SendData(w, http.StatusOK, "successfully deleted product")
}
