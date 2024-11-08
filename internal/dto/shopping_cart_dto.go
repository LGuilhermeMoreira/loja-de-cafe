package dto

type InputAddItemShoppingCartDto struct {
	CoffeeID string  `json:"coffee_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type OutputShoppingCartDto struct {
	TotalPrice   float64                      `json:"total_price"`
	CartItemList map[string]OutputCartItemDto `json:"coffee_list"`
}

type OutputCreateShoppingCartDto struct {
	Id string `json:"shopping_cart_id"`
}

type OutputCartItemDto struct {
	CoffeeID string  `json:"coffee_id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
