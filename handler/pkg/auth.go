package handler

import (
	store "github.com/cadeusept/store_project"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input store.User

	// TODO: gRPC
}

type SignInInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input SignInInput

	// TODO: gRPC
}
