package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
)

type Ticket interface {
	Create() error
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
	panic("implement me")
}

// @Summary RemoveTicket
// @Tags ticket
// @Description Remove ticket
// @ID remove-ticket
// @Accept json
// @Produce json
// @Router /api/ticket/remove [delete]
func (h *Handler) RemoveTicket(ctx *gin.Context) {
	panic("implement me")
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
