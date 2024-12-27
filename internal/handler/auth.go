package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"net/http"
)

type Auth interface {
	SignIn(user models.User) error
}

// @Summary SignIn
// @Tags auth
// @Description Login user
// @ID login
// @Accept json
// @Produce json
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(ctx *gin.Context) {
	var input models.User
	if err := ctx.ShouldBind(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	err := h.services.Auth.SignIn(input)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
