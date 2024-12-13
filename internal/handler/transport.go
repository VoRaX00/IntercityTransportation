package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
)

type Transport interface {
	Add(transport services.AddTransport) error
	Update(transport services.AddTransport) error
	Delete(stateNumber string) error
	GetAll() []models.Transport
}

// @Summary AddTransport
// @Tags transport
// @Description Add transport
// @ID add-transport
// @Accept json
// @Produce json
// @Router /api/transport/add [post]
func (h *Handler) AddTransport(ctx *gin.Context) {
	panic("implement me")
}

// @Summary UpdateTransport
// @Tags transport
// @Description Update transport
// @ID update-transport
// @Accept json
// @Produce json
// @Router /api/transport/update [put]
func (h *Handler) UpdateTransport(ctx *gin.Context) {
	panic("implement me")
}

// @Summary UpdatePartialTransport
// @Tags transport
// @Description Update partial transport
// @ID update-partial-transport
// @Accept json
// @Produce json
// @Router /api/transport/update [patch]
func (h *Handler) UpdatePartialTransport(ctx *gin.Context) {
	panic("implement me")
}

// @Summary DeleteTransport
// @Tags transport
// @ID delete-transport
// @Accept json
// @Produce json
// @Router /api/transport/delete [delete]
func (h *Handler) DeleteTransport(ctx *gin.Context) {
	panic("implement me")
}

// @Summary GetAllTransport
// @Tags transport
// @ID get-all-transport
// @Accept json
// @Produce json
// @Router /api/transport/ [get]
func (h *Handler) GetAllTransport(ctx *gin.Context) {
	panic("implement me")
}
