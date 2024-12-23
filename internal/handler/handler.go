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

	transport := api.Group("/bus")
	{
		transport.POST("/add", h.AddBus)
		transport.PUT("/update", h.UpdateBus)
		transport.PATCH("/update", h.UpdatePartialBus)
		transport.DELETE("/delete", h.DeleteBus)
		transport.GET("/", h.GetAllBus)
	}

	schedule := api.Group("/flight")
	{
		schedule.POST("/add", h.AddFlight)
		schedule.DELETE("/delete", h.DeleteFlight)
		schedule.GET("/", h.GetAllFlight)
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
