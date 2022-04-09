package domain

type Houses struct {
	ID       string `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Owner_ID string `json:"owner_id" form:"owner_id" binding:"required"`
}
