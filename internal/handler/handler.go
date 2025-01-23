package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "kursachDB/docs"
	"time"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Content-Length"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}))

	// Swagger документация
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Маршруты для аутентификации
	auth := router.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/loginAdmin", h.LoginAdmin)
	}

	// Прочие маршруты API
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
		schedule.GET("", h.GetAllFlight) // ваш маршрут для рейсов
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

	user := api.Group("/users")
	{
		user.GET("/", h.GetAllUsers)
	}

	return router
}
