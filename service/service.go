package service

import "net/http"

type ProductService interface {
	Insert(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}
