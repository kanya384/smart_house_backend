package domain

type DeviceType struct {
	ID    string `json:"id" form:"id"`
	Name  string `json:"name" form:"name" binding:"required"`
	Photo string `json:"photo" form:"photo" binding:"required"`
}
