package usecases

import (
	"go-api/app/models"
	"time"
)

func (a *APIGOUsecase) HealthCheck() models.HealthCheck {
	healthCheck := models.HealthCheck{
		Message: "OK",
		Time:    time.Now(),
	}
	return healthCheck
}
