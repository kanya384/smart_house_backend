package domain

type House struct {
	ID      string `json:"id" form:"id"`
	Name    string `json:"name" form:"name" binding:"required"`
	OwnerID string `json:"owner_id" form:"owner_id" binding:"required"`
}
