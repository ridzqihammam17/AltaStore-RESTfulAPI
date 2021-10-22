package products

type PostProductRequest struct {
	Name  string `json:"name" form:"name"`
	Price string `json:"price" form:"price"`
	Stock string `json:"stock" form:"stock"`
}

type EditProductRequest struct {
	Name  string `json:"name" form:"name"`
	Price string `json:"price" form:"price"`
	Stock string `json:"stock" form:"stock"`
}
