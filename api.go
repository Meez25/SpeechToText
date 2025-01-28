package handlers

import (
	"log/slog"
	"net/http"

	"github.com/meez25/microservice/data"
)

type Products struct {
	l *slog.Logger
}

func NewProducts(l *slog.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		p.getProducts(res, req)
		return
	}

	if req.Method == http.MethodPost {
		p.addProduct(res, req)
		return
	}

	res.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(res http.ResponseWriter, _ *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(res)
	if err != nil {
		http.Error(res, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(res http.ResponseWriter, req *http.Request) {
	p.l.Info("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(req.Body)
	if err != nil {
		http.Error(res, "Could not parse the json", http.StatusBadRequest)
	}

	p.l.Info("Prod", "Item", prod)
}
