package products

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5/pgxpool"
)

type Controller struct {
    Service *Service
}

func NewController(service *Service) *Controller {
    return &Controller{Service: service}
}

type Module struct {
    Controller *Controller
    Service    *Service
    Repository *Repository
}

func NewModule(db *pgxpool.Pool) *Module {
    repo       := NewRepository(db)
    service    := NewService(repo)
    controller := NewController(service)
    return &Module{
        Controller: controller,
        Service:    service,
        Repository: repo,
    }
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
