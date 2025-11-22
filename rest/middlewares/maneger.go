package middlewares

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddleware []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleware: make([]Middleware, 0),
	}

}
func (mngr *Manager) Use(Middlewares ...Middleware) {
	mngr.globalMiddleware = append(mngr.globalMiddleware, Middlewares...)
}

func (mngr *Manager) With(next http.Handler, Middlewares ...Middleware) http.Handler {

	n := next
	for _, Middleware := range Middlewares {
		n = Middleware(n)

	}
	// for _, globalMiddleware := range mngr.globalMiddleware {
	// 	n = globalMiddleware(n)
	// }
	return n

}

func (mngr *Manager) WrapMux(Handler http.Handler) http.Handler {

	h := Handler
	for _, Middleware := range mngr.globalMiddleware {
		h = Middleware(h)

	}
	return h

}
