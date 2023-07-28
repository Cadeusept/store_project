package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	authServices auth.Service
}

func NewHandler(authServices *auth.Service) *Handler {
	return &Handler{authServices: authServices}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/signin", h.signIn)
	}

	return router
}
