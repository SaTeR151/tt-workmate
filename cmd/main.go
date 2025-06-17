package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/sater-151/tt-workmate/internal/config"
	"github.com/sater-151/tt-workmate/internal/handlers"
	logg "github.com/sater-151/tt-workmate/internal/logger"
	"github.com/sater-151/tt-workmate/internal/service"
	logger "github.com/sirupsen/logrus"
)

func main() {
	logg.Init()
	if err := godotenv.Load(); err != nil {
		logger.Error(err)
		return
	}
	service := service.CreateTaskList()

	serverConfig := config.GetServerConfig()
	r := chi.NewRouter()

	r.Post("/api/task/new", handlers.CreateTask(service))
	r.Delete("/api/task/delete", handlers.DeleteTask(service))
	r.Get("/api/task/info", handlers.GetTaskInfo(service))

	logger.Info(fmt.Sprintf("server start at port: %s\n", serverConfig.Port))
	if err := http.ListenAndServe(":"+serverConfig.Port, r); err != nil {
		logger.Error(fmt.Sprintf("Server error: %s\n", err.Error()))
		return
	}
}
