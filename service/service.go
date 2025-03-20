package service

type ProductService interface {
	Create()
	Read(id string)
	Update(id string)
	Delete(id string)
}
