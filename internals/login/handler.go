package login

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler{
	return &Handler{
		Service : service,
	}
}

func (h *Handler) Login(c *gin.Context) {
	var req UserLogin
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		// c.JSON(http.StatusBadRequest, gin.H{
		// 	"message" : "Request not found",
		// })
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10* time.Second)
	defer cancel()

	token, err := h.Service.Login(ctx, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized,
		// 	 gin.H{
		// 	"message" : "Unauthorized",
		// }
		err.Error())
		return
	}

	c.SetCookie(
		"token",
		token,
		3600*24,
		"/",	// path ทั้งเว็บ
		"",     // domain (ว่าง = current domain)
		true,  // secure = true → ใช้กับ HTTPS เท่านั้น
		true,  // httpOnly = true → JS อ่าน cookie ไม่ได้
	)
	fmt.Println("Login success", req.Username)
	c.JSON(http.StatusOK, gin.H{
		"token" : token,
	})
}