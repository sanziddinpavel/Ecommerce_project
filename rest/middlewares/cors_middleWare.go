package middlewares

import (
	"log"
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "Application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PATCH,PUT,OPTIONS")
		w.Header().Set("Access-Control-Allow-Header", "Content-Type")
		log.Println("ami corsmiddleware 2nd e print hobo")
		next.ServeHTTP(w, r)
	})

}
