package handler
import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jackc/pgx/v5/pgxpool"
	"api/internals/user/service"
	"api/internals/user/repository"
    "api/internals/user/model"
)

type Controller struct {
    Service *service.Service
}

func NewController(service *service.Service) *Controller {
    return &Controller{Service: service}
}

type Module struct {
    Controller *Controller

}

func NewModule(db *pgxpool.Pool) *Module {
    repo   := repository.NewRepository(db)
    svc    := service.NewService(repo)
    controller := NewController(svc)
    return &Module{
        Controller: controller,
    }
}

func (ctrl *Controller) GetUser(c *gin.Context) {
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

func (ctrl *Controller) CreateUser(c *gin.Context){
    var user model.UserModel
    if err := c.ShouldBindJSON(&user);err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error" : err.Error(),
        })
        return
    }

    if err := ctrl.Service.CreateUser(c.Request.Context(),&user); err != nil{
        c.JSON(http.StatusInternalServerError, gin.H{
            "error" : err.Error(),
        })
        return
    }
    c.JSON(http.StatusCreated,gin.H{
        "id" : user.Id,
        "name" : user.Name,
    })

}