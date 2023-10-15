package server

import (
	"go-api/app/handlers"
	"go-api/app/usecases"

	"github.com/gin-gonic/gin"
)

var BaseURL = "/api"

func APIGORoutes(r *gin.Engine, us usecases.APIGOUsecaseInterface) {

	handler := &handlers.APIGOHandler{
		APIGOUsecase: us,
	}

	route := r.Group(BaseURL)
	route.GET("/health-check", handler.Check)

	user := route.Group("/users")
	user.POST("/register", handler.UserRegister)
	user.POST("/login", handler.UserLogin)
	user.PUT("/update", AuthMiddleware(), handler.UserUpdate)
	user.GET("", AuthMiddleware(), handler.GetUserDetail)
}
