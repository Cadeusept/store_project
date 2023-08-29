package controllerhttp

import (
	"net/http"
	"store-project/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create
// @Tags transactions
// @Description create new transaction
// @ID create-transaction
// @Accept json
// @Produce json
// @Param input body models.Transaction true "transaction info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /transaction/create [post]
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

// @Summary ChangeStatus
// @Tags transactions
// @Description change transaction status
// @ID changestatus-transaction
// @Accept json
// @Produce json
// @Param transaction-id path int true "transaction id"
// @Param transaction-status body updateStatusInput true "transaction status"
// @Success 200 {integer} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /transaction/changestatus/{transaction-id} [put]
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

// @Summary CheckStatus
// @Tags transactions
// @Description check transaction status
// @ID checkstatus-transaction
// @Accept json
// @Produce json
// @Param transaction-id path int true "transaction id"
// @Success 200 {object} map[string]interface{}
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /transaction/checkstatus/{transaction-id} [get]
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

// @Summary GetTransactionsByUserId
// @Tags transactions
// @Description get transactions by user id
// @ID getbyid-transactions
// @Accept json
// @Produce json
// @Param user-id path int true "user id"
// @Success 200 {iobject} getTransactionsItemsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /transaction/get/userid/{user-id} [get]
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

// @Summary GetTransactionsByUserEmail
// @Tags transactions
// @Description get transactions by user email
// @ID getbyemail-transactions
// @Accept json
// @Produce json
// @Param user-email path int true "user email"
// @Success 200 {object} getTransactionsItemsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /transaction/get/email/{user-email} [get]
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

// @Summary CancelTransaction
// @Tags transactions
// @Description cancel transaction
// @ID cancel-transaction
// @Accept json
// @Produce json
// @Param transaction-id path int true "transaction id"
// @Success 200 {integer} statusResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /transaction/cancel/{transaction-id} [post]
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
