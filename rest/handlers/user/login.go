package user

import (
	"Ecommerce/config"
	"Ecommerce/database"
	"Ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "invalid request data", http.StatusBadRequest)
		return
	}
	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil {
		http.Error(w, "invalid ", http.StatusBadRequest)
		return
	}

	cnf := config.GetConfig()
	accessToken, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})
	if err != nil {
		http.Error(w, "internal server Error", http.StatusInternalServerError)
		return
	}
	util.SendData(w, accessToken, http.StatusCreated)
}
