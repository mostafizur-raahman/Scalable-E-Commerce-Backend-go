package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddleware []Middleware
}

func New() *Manager {
	return &Manager{
		globalMiddleware: make([]Middleware, 0),
	}
}

func (m *Manager) Use(middlewares ...Middleware) {
	m.globalMiddleware = append(m.globalMiddleware, middlewares...)
}

func (m *Manager) Chain(next http.Handler, middlewares ...Middleware) http.Handler {
	all := append(m.globalMiddleware, middlewares...)

	for i := range m.globalMiddleware {
		next = all[i](next)
	}

	return next
}
