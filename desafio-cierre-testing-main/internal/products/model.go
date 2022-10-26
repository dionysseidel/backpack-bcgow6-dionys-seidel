package products

type Product struct {
	ID          string  `json:"ID"`
	SellerID    string  `json:"SellerID"`
	Description string  `json:"Description"`
	Price       float64 `json:"Price"`
}
