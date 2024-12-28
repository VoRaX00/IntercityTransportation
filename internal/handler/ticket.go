package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"kursachDB/internal/services/ticket"
	"net/http"
	"strconv"
)

type Ticket interface {
	BuyTicket(ticket services.BuyTicket) error
	RemoveTicket(id int64) error
	GetAll() ([]models.Ticket, error)
	GetByUser(userId int64) ([]models.Ticket, error)
}

// @Summary BuyTicket
// @Tags ticket
// @Description Buy ticket
// @ID buy-ticket
// @Accept json
// @Produce json
// @Param input body services.BuyTicket true "Ticket info for add"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/ticket/buy [post]
func (h *Handler) BuyTicket(ctx *gin.Context) {
	var input services.BuyTicket
	if err := ctx.ShouldBind(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err := h.services.Ticket.BuyTicket(input)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// @Summary RemoveTicket
// @Tags ticket
// @Description Remove ticket
// @ID remove-ticket
// @Accept json
// @Produce json
// @Param id path int64 true "Id of the ticket"
// @Router /api/ticket/{id} [delete]
func (h *Handler) RemoveTicket(ctx *gin.Context) {
	param, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	id := int64(param)
	if err = h.services.Ticket.RemoveTicket(id); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

// @Summary GetAllTickets
// @Tags ticket
// @Description Get all tickets
// @ID get-all-tickets
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/ticket [get]
func (h *Handler) GetAllTickets(ctx *gin.Context) {
	res, err := h.services.Ticket.GetAll()
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// @Summary GetUserTickets
// @Tags ticket
// @Description Get user tickets
// @ID get-user-tickets
// @Accept json
// @Produce json
// @Param phoneNumber query int64 false "Phone number of the user"
// @Success 200 {array} models.Ticket
// @Failure 500 {object} map[string]string
// @Router /api/ticket/user [get]
func (h *Handler) GetUserTickets(ctx *gin.Context) {
	param, err := strconv.Atoi(ctx.Query("phoneNumber"))
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	phoneNumber := int64(param)
	tickets, err := h.services.Ticket.GetByUser(phoneNumber)
	if err != nil {
		if errors.Is(err, ticket.ErrNotFound) {
			responses.NewErrorResponse(ctx, http.StatusNotFound, ErrRecordNotFound)
			return
		}

		responses.NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, tickets)
}
