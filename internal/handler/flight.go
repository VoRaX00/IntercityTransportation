package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"kursachDB/internal/domain/models"
	"kursachDB/internal/handler/responses"
	"kursachDB/internal/services"
	"kursachDB/internal/services/flight"
	"net/http"
	"strconv"
	"time"
)

const (
	ErrFlightNotFound = "Flight not found"
)

type Flight interface {
	Add(flight services.AddFlight) error
	Delete(id int) error
	GetAll(filters services.FlightFilter) ([]models.Flight, error)
}

// @Summary AddFlight
// @Tags flight
// @Description Add new flight
// @ID add-flight
// @Accept json
// @Produce json
// @Param input body services.AddFlight true "Flight info for add"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/flight/add [post]
func (h *Handler) AddFlight(ctx *gin.Context) {
	var input services.AddFlight
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	if err := validateAddFlight(input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Flight.Add(input); err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

func validateAddFlight(flight services.AddFlight) error {
	departure, err := time.Parse("02.01.2006", flight.Departure)
	if err != nil {
		return fmt.Errorf("departure date format is wrong")
	}

	arrival, err := time.Parse("02.01.2006", flight.Arrival)
	if err != nil {
		return fmt.Errorf("arrival date format is wrong")
	}

	if departure.After(arrival) {
		return fmt.Errorf("departure date is wrong")
	}
	return nil
}

// @Summary DeleteFlight
// @Tags flight
// @Description Delete flight
// @ID delete-flight
// @Accept json
// @Produce json
// @Param id path int64 true "Id of the flight"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/flight/{id} [delete]
func (h *Handler) DeleteFlight(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, ErrInvalidArguments)
		return
	}

	if err = h.services.Flight.Delete(id); err != nil {
		if errors.Is(err, flight.ErrFlightNotFound) {
			responses.NewErrorResponse(ctx, http.StatusNotFound, ErrFlightNotFound)
			return
		}

		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"status": "success",
		},
	)
}

// @Summary GetAllFlight
// @Tags flight
// @Description Get all flight
// @ID get-all-flight
// @Accept json
// @Produce json
// @Param from query string false "Name of the from"
// @Param to query string false "Name of the to"
// @Param stateNumber query string false "State number of the bus"
// @Success 200 {array} models.Flight
// @Failure 500 {object} map[string]string
// @Router /api/flight [get]
func (h *Handler) GetAllFlight(ctx *gin.Context) {
	from := ctx.Query("from")
	to := ctx.Query("to")
	stateNumber := ctx.Query("stateNumber")

	filters := services.FlightFilter{
		From:        from,
		To:          to,
		StateNumber: stateNumber,
	}

	res, err := h.services.Flight.GetAll(filters)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusInternalServerError, ErrInternalServer)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
