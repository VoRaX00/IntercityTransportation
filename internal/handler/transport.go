package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"net/http"
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
	var input services.AddTransport
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	if err := h.services.Transport.Add(input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInvalidArguments)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

func validateAddTransport(transport services.AddTransport) error {
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
	stateNumber := ctx.Query("stateNumber")
	if err := h.services.Transport.Delete(stateNumber); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInvalidArguments)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

// @Summary GetAllTransport
// @Tags transport
// @ID get-all-transport
// @Accept json
// @Produce json
// @Router /api/transport/ [get]
func (h *Handler) GetAllTransport(ctx *gin.Context) {
	res := h.services.Transport.GetAll()
	ctx.JSON(http.StatusOK, res)
}
