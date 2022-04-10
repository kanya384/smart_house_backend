package domain

type Device struct {
	ID           string `json:"id" form:"id"`
	DeviceTypeId string `json:"device_type_id" form:"device_type_id" binding:"required"`
	HousePartId  string `json:"house_part_id" form:"house_part_id" binding:"required"`
}
