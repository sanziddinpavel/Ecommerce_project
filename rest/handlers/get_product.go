package handlers

import (
	"Ecommerce/database"
	"Ecommerce/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product_id", 400)
		return
	}
	product := database.Get(pid)
	if product == nil {
		util.SendError(w, 404, "product not found")
		return

	}

	util.SendData(w, product, 200)
}
