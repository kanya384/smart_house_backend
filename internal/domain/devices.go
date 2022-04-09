package domain

type Devices struct {
	ID           string `json:"id" form:"id"`
	DeviceTypeId string `json:"device_type_id" form:"device_type_id" binding:"required"`
}
