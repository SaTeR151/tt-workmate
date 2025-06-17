package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sater-151/tt-workmate/internal/service"
	logger "github.com/sirupsen/logrus"
)

type RespoceId struct {
	Id string `json:"id"`
}

func CreateTask(s service.ServiceInterface) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Info("creating task")
		id := s.CreateTask()
		logger.Info("task created")
		go s.Processing(id)
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(RespoceId{Id: id})
		logger.Info("id sended")
	}
}

func DeleteTask(s service.ServiceInterface) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Info("deleting task")
		id := req.FormValue("id")
		if id == "" {
			logger.Error("id required")
			http.Error(res, "id required", http.StatusBadRequest)
			return
		}
		err := s.DeleteTask(id)
		if err != nil {
			logger.Error(err.Error())
			if err == service.ErrorTaskNotFound {
				http.Error(res, "Error: "+err.Error(), http.StatusBadRequest)
			} else {
				http.Error(res, "Error: "+err.Error(), http.StatusForbidden)
			}
			return
		}
		res.WriteHeader(http.StatusNoContent)
	}
}

func GetTaskInfo(s service.ServiceInterface) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Info("getting task info")
		id := req.FormValue("id")
		if id == "" {
			logger.Error("id required")
			http.Error(res, "id required", http.StatusBadRequest)
			return
		}
		taskInfo, err := s.GetTaskInfo(id)
		if err != nil {
			logger.Error(err.Error())
			http.Error(res, "Error: "+err.Error(), http.StatusBadRequest)
			return
		}
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(taskInfo)
		logger.Info("task info send")
	}
}
