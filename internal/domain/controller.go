package domain

type Controller struct {
	ID               string `json:"id" form:"id"`
	ControllerTypeId string `json:"controller_type_id" form:"controller_type_id" binding:"required"`
	Ip               string `json:"ip" form:"ip" binding:"required"`
}
