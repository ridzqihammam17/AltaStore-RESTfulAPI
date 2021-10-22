package products

type GetProductResponse struct {
	Name  string `json:"name" form:"name"`
	Price string `json:"price" form:"price"`
	Stock string `json:"stock" form:"stock"`
}
