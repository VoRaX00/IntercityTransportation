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

func (h *Handler) GetAllUsers(ctx *gin.Context) {
	res, err := h.services.User.GetAll()
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	ctx.JSON(200, res)
}
