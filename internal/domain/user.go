package domain

type User struct {
	ID      string `json:"id" form:"id"`
	Name    string `json:"name" form:"name" binding:"required"`
	Surname string `json:"surname" form:"surname" binding:"required"`
}
