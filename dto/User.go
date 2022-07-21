package dto

type User struct {
	Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"-" form:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
