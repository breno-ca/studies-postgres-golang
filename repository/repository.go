package repository

import "studies-postgres-golang/model"

type ProductRepository interface {
	Insert(product model.Product) error
	Index() ([]model.Product, error)
}
