package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"net/http"
)

type Auth interface {
	SignIn(login services.UserLogin) error
}

// @Summary SignIn
// @Tags auth
// @Description Login user
// @ID login
// @Accept json
// @Produce json
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(ctx *gin.Context) {
	var input services.UserLogin
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
