package handlers

import (
	"Ecommerce/database"
	"Ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "give me valid json", 400)
		return
	}
	createdProduct := database.Store(newProduct)

	util.SendData(w, createdProduct, 201)
	fmt.Println(createdProduct)

}
