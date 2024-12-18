package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api")

	transport := api.Group("/transactions")
	{
		transport.POST("/add", h.AddTransport)
		transport.PUT("/update", h.UpdateTransport)
		transport.PATCH("/update", h.UpdatePartialTransport)
		transport.DELETE("/delete", h.DeleteTransport)
		transport.GET("/", h.GetAllTransport)
	}

	schedule := api.Group("/schedule")
	{
		schedule.POST("/add", h.AddSchedule)
		schedule.DELETE("/delete", h.DeleteSchedule)
		schedule.GET("/", h.GetAllSchedule)
	}

	place := api.Group("/place")
	{
		place.POST("/add", h.AddPlace)
		place.DELETE("/delete", h.DeletePlace)
		place.GET("/", h.GetAllPlaces)
	}

	ticket := api.Group("/ticket")
	{
		ticket.POST("/buy", h.BuyTicket)
		ticket.DELETE("remove-ticket", h.RemoveTicket)
		ticket.GET("/", h.GetAllTickets)
	}
	return router
}
