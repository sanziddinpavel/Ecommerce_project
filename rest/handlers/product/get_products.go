package product

import (
	"Ecommerce/database"
	"Ecommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {

	//  handlePreflightReq(w,r)
	util.SendData(w, database.List(), 200)

}
