package controllerhttp

import (
	"net/http"
	"store-project/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) create(c *gin.Context) {
	var input models.Transaction

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.uc.Create(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type updateStatusInput struct {
	NewStatus string `json:"newstatus" binding:"required"`
}

func (h *Handler) changeStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input updateStatusInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.uc.ChangeStatus(id, input.NewStatus)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) checkStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	currStatus, err := h.uc.CheckStatusById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": currStatus,
	})
}

type getTransactionsItemsResponse struct {
	Data []models.Transaction `json:"data"`
}

func (h *Handler) getByUserId(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	items, err := h.uc.GetTransactionsByUserId(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getTransactionsItemsResponse{
		Data: items,
	})
}

func (h *Handler) getByUserEmail(c *gin.Context) {
	email := c.Param("email")

	if err := validateEmail(email); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid email")
		return
	}

	items, err := h.uc.GetTransactionsByUserEmail(email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getTransactionsItemsResponse{
		Data: items,
	})
}

func (h *Handler) cancelById(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.uc.CancelTransactionById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
