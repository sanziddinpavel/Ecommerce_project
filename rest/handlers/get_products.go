package handlers

import (
	"Ecommerce/database"
	"Ecommerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	//  handlePreflightReq(w,r)
	util.SendData(w, database.List(), 200)

}
