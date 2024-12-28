package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"net/http"
)

type Auth interface {
	Login(user models.User) error
}

// @Summary Login
// @Tags auth
// @Description Login user
// @ID login
// @Accept json
// @Produce json
// @Param input body models.User true "User info for login"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/login [post]
func (h *Handler) Login(ctx *gin.Context) {
	var input models.User
	if err := ctx.ShouldBind(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err := h.services.Auth.Login(input)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
