package login

type UserLogin struct {
	Username string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
