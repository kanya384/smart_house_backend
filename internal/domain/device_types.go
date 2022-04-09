package domain

type DeviceTypes struct {
	ID   string `json:"id" form:"id"`
	Name string `json:"name" form:"name" binding:"required"`
}
