package dto

type Book struct {
	Model
	Name        string `json:"name" form:"name" inding:"required"`
	Author      string `json:"author" form:"author" inding:"required"`
	Description string `json:"description" form:"description"`
	Publisher   string `json:"publisher" form:"publisher" inding:"required"`
	ISBN        string `json:"isbn" form:"isbn" binding:"required"`
}
