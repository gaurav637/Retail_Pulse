package routes

import (
	"log"

	"github.com/kiranaClub/config"
	"github.com/kiranaClub/internal/database"
	"github.com/kiranaClub/internal/handler"
	"github.com/kiranaClub/internal/repository"
	"github.com/kiranaClub/internal/services"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	cfg := config.LoadConfig()
	db, err := database.DBConnecting(cfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Job setup
	jobRepository := repository.NewJobRepository(db)
	jobService := services.NewJobService(jobRepository)
	jobHandler := handler.NewJobHandler(jobService)

	api := e.Group("/api")

	// job routes
	api.POST("/submit", jobHandler.InsertJob)
	api.GET("/status/job", jobHandler.GetJobByID)

}
