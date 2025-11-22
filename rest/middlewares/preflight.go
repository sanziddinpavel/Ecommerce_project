package middlewares

import (
	"log"
	"net/http"
)

func PreflightReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		log.Println("ami preflight 3rd e print hobo")
		next.ServeHTTP(w, r)
	})

}
