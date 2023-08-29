package controllerhttp

import (
	"store-project/internal/handler/usecase"

	_ "store-project/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	uc usecase.HandlerUseCaseI
}

func NewHandler(uc usecase.HandlerUseCaseI) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	transaction := router.Group("/transaction")
	{
		transaction.POST("/create", h.create)
		transaction.PUT("/changestatus/:id", h.changeStatus)
		transaction.GET("/checkstatus/:id", h.checkStatus)

		get := transaction.Group("/get")
		{
			get.GET("/userid/:id", h.getByUserId)
			get.GET("/email/:email", h.getByUserEmail)
		}

		transaction.POST("/cancel/:id", h.cancelById)
	}

	return router
}
