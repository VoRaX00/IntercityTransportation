package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	services *Service
}

func NewHandler(services *Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {})
	}
	return router
}
