package handler

import (
	"github.com/gin-gonic/gin"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"net/http"
)

type Auth interface {
	SignIn(login services.UserLogin) (services.Tokens, error)
	SignUp(register services.UserRegister) (int, error)
	RefreshTokens(tokens services.Tokens) (services.Tokens, error)
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

	tokens, err := h.services.Auth.SignIn(input)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
	}

	ctx.JSON(http.StatusOK, tokens)
}

// @Summary SignUp
// @Tags auth
// @Description Registration user
// @ID registration
// @Accept json
// @Produce json
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(ctx *gin.Context) {
	var input services.UserRegister
	if err := ctx.ShouldBind(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	id, err := h.services.Auth.SignUp(input)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"id": id,
		},
	)
}

// @Summary RefreshToken
// @Tags auth
// @Description Refresh tokens
// @ID refresh-tokens
// @Accept json
// @Produce json
// @Router /auth/refresh-tokens [put]
func (h *Handler) RefreshTokens(ctx *gin.Context) {
	var input services.Tokens
	if err := ctx.ShouldBind(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	tokens, err := h.services.Auth.RefreshTokens(input)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}
