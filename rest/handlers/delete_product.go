package handlers

import (
	"Ecommerce/database"
	"Ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product_id", 400)
		return

	}
	database.Delete(pid)
	util.SendData(w, "successfully deleted product", 201)
}
