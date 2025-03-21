package product

import "database/sql"

const (
	// Error Messages
	errPreparingStmt = "Error preparing statement: "
	errExecutingStmt = "Error executing statement: "
	errFetchingRows  = "Error fetching rows: "
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}
