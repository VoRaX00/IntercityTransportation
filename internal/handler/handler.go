package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "kursachDB/docs"
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
		auth.POST("/login", h.Login)
	}

	api := router.Group("/api")

	transport := api.Group("/bus")
	{
		transport.POST("/add", h.AddBus)
		transport.GET("/", h.GetAllBus)
		transport.GET("/:stateNumber", h.GetBus)
	}

	schedule := api.Group("/flight")
	{
		schedule.POST("/add", h.AddFlight)
		schedule.GET("/", h.GetAllFlight)
	}

	place := api.Group("/place")
	{
		place.POST("/add", h.AddPlace)
		place.GET("/", h.GetAllPlaces)
	}

	ticket := api.Group("/ticket")
	{
		ticket.POST("/buy", h.BuyTicket)
		ticket.DELETE("/:id", h.RemoveTicket)
		ticket.GET("/", h.GetAllTickets)
		ticket.GET("/user", h.GetUserTickets)
	}

	user := api.Group("/user")
	{
		user.GET("/", h.GetAllUsers)
	}
	return router
}
