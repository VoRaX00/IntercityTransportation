package mapper

import (
	"kursachDB/internal/domain/models"
	"kursachDB/internal/services"
)

func PlaceAddToPlace(place services.AddPlace) models.Place {
	return models.Place{
		NamePlace: place.Name,
		Type: models.TypePlace{
			Type: place.Type,
		},
	}
}
