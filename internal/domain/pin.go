package domain

type Pin struct {
	ID           string `json:"id" form:"id"`
	ControllerId string `json:"controller_id" form:"controller_id" binding:"required"`
	DeviceId     string `json:"device_id" form:"device_id" binding:"required"`
	Value        int    `json:"value" form:"value" binding:"required"`
	Type         int    `json:"type" form:"type" binding:"required"` // 0 - digital, 1 - analog
}
