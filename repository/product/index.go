package product

import (
	"errors"

	"studies-postgres-golang/model"
)

func (r *ProductRepository) Index() ([]model.Product, error) {
	rows, err := r.db.Query(`

		SELECT 
			id, name, price, available
		FROM 
			product

	`)
	if err != nil {
		return nil, errors.New(errFetchingRows + err.Error())
	}
	defer rows.Close()

	var products []model.Product

	for rows.Next() {

		var product model.Product

		err = rows.Scan(
			&product.ID, &product.Name, &product.Price, &product.Available,
		)
		if err != nil {
			return nil, errors.New(errFetchingRows + err.Error())
		}

		products = append(products, product)
	}

	return products, nil
}
