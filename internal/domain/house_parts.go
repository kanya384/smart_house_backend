package domain

type HouseParts struct {
	ID      string `json:"id" form:"id"`
	Name    string `json:"name" form:"name" binding:"required"`
	HouseID string `json:"house_id" form:"house_id" binding:"required"`
}
