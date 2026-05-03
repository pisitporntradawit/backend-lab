package handler

import (
	"context"
	"log"
	"net/http"
	"time"
	"api/internals/user/model"
	"api/internals/user/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}

func (ctrl *Handler) CreateUser(c *gin.Context) {
	var user model.UserModel
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := ctrl.Service.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":   user.Id,
		"name": user.Name,
	})

}

func (ctrl *Handler) GetUser(c *gin.Context) {
	resultUser, err := ctrl.Service.Getuser(c.Request.Context())
	if err != nil {
		log.Printf("GetProducts error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve user",
		})
		return
	}
	c.JSON(http.StatusOK, resultUser)
}

func (ctrl *Handler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID not found",
		})
		return
	}
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user, err := ctrl.Service.GetUserByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   user.Id,
		"name": user.Name,
	})
}

func (ctrl *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "connot find id",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := ctrl.Service.DeleteUser(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "user.handler.delete",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "delete user success",
	})
}
