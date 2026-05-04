package route

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Interface แทนการผูกกับ concrete struct
type userInterface interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type productInterface interface {
	GetProducts(c *gin.Context)
}

type loginInterface interface {
	 Login(c *gin.Context)
}

func RouterAPI(user userInterface, product productInterface, login loginInterface) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(corsConfig()))
	// r.GET("/users", user.GetUser)
	// r.POST("/users/newUser", user.CreateUser)
	// r.GET("/users/:id", user.GetUserByID)
	// r.DELETE("/users/:id", user.DeleteUser)
	r.POST("/login", login.Login)
	r.GET("/productss", product.GetProducts)
	v1 := r.Group("/api")
    groupUser(v1, user)
	return r
}

func corsConfig() cors.Config {
	origins := os.Getenv("ALLOWED_ORIGINS")
	allowOrigins := []string{"http://localhost:3000"}
	if origins != "" {
		allowOrigins = strings.Split(origins, ",")
	}

	return cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
}

// func groupProducts(rg *gin.RouterGroup, ctrl userInterface) {
// 	products := rg.Group("/products")
// 	{
// 		products.GET("", ctrl.GetUser)
// 	}
// }

func groupUser(rg *gin.RouterGroup, userInf userInterface) {
	user := rg.Group("/users")
	{
		user.GET("", userInf.GetUser)
		user.POST("", userInf.CreateUser)
		user.GET("/:id", userInf.GetUserByID)
		user.DELETE("/:id", userInf.DeleteUser)
	}
}
