package handlers

import (
	"encoding/json"
	"gomic/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	encoder := json.NewEncoder(rw)
	encoder.Encode(lp)
}
