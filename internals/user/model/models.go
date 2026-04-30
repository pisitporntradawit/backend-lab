package model

import (
	"github.com/google/uuid"
)


type UserModel struct{
	Id uuid.UUID `json:"id"`
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
