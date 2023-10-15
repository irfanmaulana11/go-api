package usecases

import (
	"go-api/app/models"
	"go-api/app/repositories"
	"go-api/helpers"

	"github.com/gin-gonic/gin"
)

type APIGOUsecase struct {
	DBRepository repositories.DBRepositoryInterface
}

type APIGOUsecaseInterface interface {
	HealthCheck() models.HealthCheck
	UserRegister(c *gin.Context, user models.User) (*helpers.ValidationErrors, error)
	UserLogin(c *gin.Context, user models.UserLogin) (*string, error)
	GetUserDetail(c *gin.Context, email string) (*models.User, error)
	UserUpdate(c *gin.Context, email string, user models.User) (*helpers.ValidationErrors, error)
}

func NewAPIGOUsecase(dbRepo repositories.DBRepositoryInterface) APIGOUsecaseInterface {
	return &APIGOUsecase{
		DBRepository: dbRepo,
	}
}
