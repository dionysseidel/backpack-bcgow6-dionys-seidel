package domain

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"nombre" binding:"required"`
	Type        string  `json:"tipo"`
	Count       int     `json:"cantidad"`
	Price       float64 `json:"precio"`
	WarehouseId int     `json:"warehouse_id" binding:"required"`
}
