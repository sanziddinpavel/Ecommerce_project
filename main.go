package main

import (
	"Ecommerce/cmd"
	"Ecommerce/util"
	"fmt"
)

func main() {
	cmd.Server()

	jwt, err := util.CreateJwt("my-secret", util.Payload{
		Sub:         45,
		FirstName:   "Sanzid",
		LastName:    "Din Pavel",
		Email:       "pavellabib912@gmail.com",
		IsShopOwner: false,
	})
	if err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println(jwt)
}
