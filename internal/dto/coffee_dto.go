package dto

type OutputCoffee struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Id          string  `json:"id"`
	Description string  `json:"description"`
	Data        string  `json:"data"` // base64
}

type InputCreateCoffee struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	LabelFile   string  `json:"label_file"`
	Data        string  `json:"data"`
}

type InputUpdateCoffee struct {
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Data        string  `json:"data"`
}
