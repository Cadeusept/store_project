package controllerhttp

import (
	"store-project/internal/handler/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.HandlerUseCaseI
}

func NewHandler(uc usecase.HandlerUseCaseI) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	transaction := router.Group("/transaction")
	{
		transaction.POST("/create", h.create)
		transaction.PUT("/changestatus/:id", h.changeStatus)
		transaction.GET("/checkstatus/:id", h.checkStatus)

		get := transaction.Group("/get")
		{
			get.GET("?id=:id", h.getByUserId)
			get.GET("?email=:email", h.getByUserEmail)
		}

		transaction.DELETE("/:id", h.cancelById)
	}

	return router
}
