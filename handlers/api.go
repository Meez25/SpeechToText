package handlers

import (
	"log/slog"
	"net/http"
)

type Products struct {
	l *slog.Logger
}

func NewProducts(l *slog.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		return
	}

	if req.Method == http.MethodPost {
		return
	}

	res.WriteHeader(http.StatusMethodNotAllowed)
}
