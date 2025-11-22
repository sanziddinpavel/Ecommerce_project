package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Println("ami 1st logger middleware")
		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))

	})

}
