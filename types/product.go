package types

type Product struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

func (p *Product) ValidateProduct() map[string]string {
	errors := map[string]string{}
	if p.Price == 0 {
		errors["price"] = "price should not be empty"
	}
	if p.Image == "" {
		errors["image"] = "the product should have one image"
	}
	if p.Title == "" {
		errors["title"] = "insert a product title"
	}
	if p.Description == "" {
		errors["description"] = "insert a product description"
	}
	return errors
}
