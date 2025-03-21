package product

import (
	"errors"

	"studies-postgres-golang/model"
)

func (r *ProductRepository) Insert(product model.Product) error {
	stmt, err := r.db.Prepare(`

		INSERT INTO
		    product (id, name, price, available)
		VALUES
		    ($1, $2, $3, $4)

	`)
	if err != nil {
		return errors.New(errPreparingStmt + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		product.ID, product.Name, product.Price, product.Available,
	)
	if err != nil {
		return errors.New(errExecutingStmt + err.Error())
	}

	return nil
}
