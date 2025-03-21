package product

import (
	"encoding/json"
	"net/http"

	"studies-postgres-golang/model"

	"github.com/google/uuid"
)

func (uc ProductService) Insert(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product.ID = uuid.New().String()

	err = uc.repository.Insert(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
