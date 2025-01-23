package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"net/http"
)

type Users interface {
	GetAll() ([]models.User, error)
}

// @Summary GetAllUsers
// @Tags users
// @Description Get all users
// @ID users
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/users [get]
func (h *Handler) GetAllUsers(ctx *gin.Context) {
	res, err := h.services.User.GetAll()
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInternalServer)
	}
	ctx.JSON(200, res)
}
