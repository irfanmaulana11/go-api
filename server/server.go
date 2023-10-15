package server

import (
	"fmt"
	"log"
	"os"
	"strconv"

	dbr "go-api/app/repositories/dbrepositories"
	"go-api/app/usecases"
	"go-api/config"

	"github.com/gin-gonic/gin"
)

type RESTServer struct {
	router *gin.Engine
}

type RestRunner interface {
	Run()
}

func (rs *RESTServer) Run() {
	servicePort := os.Getenv("SERVICE_PORT")
	log.Println(fmt.Sprintf("Running Server in :%s", servicePort))

	if err := rs.router.Run(":" + servicePort); err != nil {
		log.Fatal(err)
	}
}

func NewRestServer() RestRunner {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware)

	dbPort, _ := strconv.ParseUint(os.Getenv("DB_PORT"), 10, 32)

	db, err := dbr.NewPostgreSQLConn(config.PostgreSQLConfiguration{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     int(dbPort),
		DBName:     os.Getenv("DB_NAME"),
		DBOptions:  os.Getenv("DB_OPTIONS"),
		Locale:     os.Getenv("DB_LOCALE"),
	})

	if err != nil {
		log.Println("Error connecting to PostgreSQL:", err)
		os.Exit(1)
	}

	dbRepo := dbr.NewDBRepository(db)

	us := usecases.NewAPIGOUsecase(dbRepo)

	APIGORoutes(r, us)

	return &RESTServer{
		router: r,
	}
}
