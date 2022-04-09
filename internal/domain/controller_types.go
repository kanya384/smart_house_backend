package domain

type ControllerTypes struct {
	ID            string `json:"id" form:"id"`
	Name          string `json:"name" form:"name" binding:"required"`
	Photo         string `json:"photo" form:"photo"`
	DigitalPinCnt int    `json:"digital_pin_cnt" form:"digital_pin_cnt" binding:"required"`
	AnalogPinCnt  int    `json:"analog_pin_cnt" form:"analog_pin_cnt" binding:"required"`
}
