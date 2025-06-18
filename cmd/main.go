package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	_ "github.com/sater-151/tt-workmate/docs"
	"github.com/sater-151/tt-workmate/internal/config"
	"github.com/sater-151/tt-workmate/internal/controller/rest"
	logg "github.com/sater-151/tt-workmate/internal/logger"
	"github.com/sater-151/tt-workmate/internal/services/taskManager"
	logger "github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Test Task I/O bound
//	@version		0.6.1
//	@description	Server for create, read and delete tasks

// @host		localhost:8080
// @BasePath	/api
func main() {
	logg.Init()
	if err := godotenv.Load(); err != nil {
		logger.Error(err)
		return
	}
	service := taskManager.New()

	serverConfig := config.GetServerConfig()
	r := chi.NewRouter()

	r.Post("/api/task/new", rest.CreateTask(service))
	r.Delete("/api/task/delete", rest.DeleteTask(service))
	r.Get("/api/task/info", rest.GetTaskInfo(service))

	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("/swagger/doc.json")))

	logger.Info(fmt.Sprintf("server start at port: %s\n", serverConfig.Port))
	if err := http.ListenAndServe(":"+serverConfig.Port, r); err != nil {
		logger.Error(fmt.Sprintf("Server error: %s\n", err.Error()))
		return
	}
}
