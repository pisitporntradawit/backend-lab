package products

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"

)

type Controller struct {
    Service *Service
}

func NewController(service *Service) *Controller {
    return &Controller{Service: service}
}

func (ctrl *Controller) GetProducts(c *gin.Context) {
    products, err := ctrl.Service.GetProducts(c.Request.Context())
    if err != nil {
        log.Printf("GetProducts error: %v", err)
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "failed to retrieve products",
        })
        return
    }
    c.JSON(http.StatusOK, products)
}
