package product

import (
	"fmt"
	"net/http"
)

func (h *Handler) Hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HELLO WORLD")
}
