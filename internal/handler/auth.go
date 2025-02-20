package handler

import (
	"github.com/gin-gonic/gin"
	_ "github.com/vektra/mockery"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/pkg/jwt"
	"net/http"
	"time"
)

//go:generate mockery --name=Auth --output=./mocks --case=underscore
type Auth interface {
	Login(user models.User) (string, error)
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

	token, err := h.services.Auth.Login(input)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}

// @Summary LoginAdmin
// @Tags auth
// @Description Login admin
// @ID login-admin
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/loginAdmin [post]
func (h *Handler) LoginAdmin(ctx *gin.Context) {
	token, err := jwt.NewToken(models.User{FIO: "admin"}, time.Hour*24)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}
