package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"net/http"
	"strconv"
)

type Ticket interface {
	BuyTicket(ticket services.BuyTicket) error
	RemoveTicket(id int64) error
	Update() error
	GetAll() ([]models.Ticket, error)
	GetByUser() ([]models.Ticket, error)
}

// @Summary BuyTicket
// @Tags ticket
// @Description Buy ticket
// @ID buy-ticket
// @Accept json
// @Produce json
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
// @Router /api/ticket/remove [delete]
func (h *Handler) RemoveTicket(ctx *gin.Context) {
	param, err := strconv.Atoi(ctx.Query("id"))
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
// @Router /api/ticket/ [get]
func (h *Handler) GetAllTickets(ctx *gin.Context) {
	panic("implement me")
}
