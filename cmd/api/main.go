package main

import (
	"log"

	config "github.com/AlexRijikov/go-activity-service/configs"
	"github.com/AlexRijikov/go-activity-service/internal/handler"
	"github.com/AlexRijikov/go-activity-service/internal/repository"
	"github.com/AlexRijikov/go-activity-service/internal/scheduler"
	"github.com/AlexRijikov/go-activity-service/internal/usecase"
	"github.com/AlexRijikov/go-activity-service/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()
	db := database.NewPostgres(cfg)

	eventRepo := repository.NewEventRepository(db)
	eventUsecase := usecase.NewEventUsecase(eventRepo)
	h := handler.NewHandler(eventUsecase)

	// cron
	sched := scheduler.NewScheduler(eventRepo)
	sched.Start()

	r := gin.Default()
	handler.SetupRoutes(r, h)

	log.Println("Server started on port:", cfg.AppPort)
	r.Run(":" + cfg.AppPort)
}
