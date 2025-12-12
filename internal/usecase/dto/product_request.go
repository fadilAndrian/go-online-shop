package dto

type CreateProductRequest struct {
	Name  string `validate:"required,min=3"`
	Price int    `validate:"required,gt=0"`
}

type UpdateProductRequest struct {
	Name  string `validate:"required,min=3"`
	Price int    `validate:"required,gt=0"`
}
