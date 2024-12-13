package handler

import "github.com/gin-gonic/gin"

// @Summary SignIn
// @Tags auth
// @Description Login user
// @ID login
// @Accept json
// @Produce json
// @Router /auth/sign-in [post]
func (h *Handler) SignIn(ctx *gin.Context) {
	panic("implement me")
}

// @Summary SignUp
// @Tags auth
// @Description Registration user
// @ID registration
// @Accept json
// @Produce json
// @Router /auth/sign-up [post]
func (h *Handler) SignUp(ctx *gin.Context) {
	panic("implement me")
}

// @Summary RefreshToken
// @Tags auth
// @Description Refresh tokens
// @ID refresh-tokens
// @Accept json
// @Produce json
// @Router /auth/refresh-tokens [put]
func (h *Handler) RefreshTokens(ctx *gin.Context) {
	panic("implement me")
}
