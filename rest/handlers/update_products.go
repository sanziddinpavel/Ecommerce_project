package handlers

import (
	"Ecommerce/database"
	"Ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pid, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valid product_id", 400)
		return

		// 	util.SendData(w, database.List(), 200)

	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "give me valid json", 400)
		return

	}

	newProduct.ID = pid
	database.Update(newProduct)
	util.SendData(w, "successfully updated product", 201)
}
