package rest

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sater-151/tt-workmate/internal/controller/rest/restutils"
	"github.com/sater-151/tt-workmate/internal/services/taskManager"
	logger "github.com/sirupsen/logrus"
)

var ErrorIDRequired = errors.New("id required")

type HTTPError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type ResponceId struct {
	Id string `json:"id"`
}

// CreateTask godoc
//
//	@Summary		Create a new task
//	@Tags			Task
//	@Produce		json
//	@Success		201	{object}	rest.ResponceId
//	@Router			/api/task/new [post]
func CreateTask(s taskManager.ServiceInterface) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Info("creating task")
		id := s.CreateTask()
		logger.Info("task created")
		go s.Processing(id)
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(ResponceId{Id: id})
		logger.Info("id sended")
	}
}

// DeleteTask godoc
//
//	@Summary		Delete task
//	@Description	Delete finished task by id
//	@Tags			Task
//	@Produce		json
//	@Param			id	query		string	true	"task id"
//	@Success		204	{object}	rest.ResponceId
//	@Failure		400	{object}	restutils.HTTPError
//	@Failure		403	{object}	restutils.HTTPError
//	@Router			/api/task/delete [delete]
func DeleteTask(s taskManager.ServiceInterface) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Info("deleting task")
		id := req.FormValue("id")
		if id == "" {
			logger.Error(ErrorIDRequired)
			restutils.Error(res, ErrorIDRequired.Error(), http.StatusBadRequest)
			return
		}
		if err := s.DeleteTask(id); err != nil {
			logger.Warn(err.Error())
			if err == taskManager.ErrorTaskNotFound {
				restutils.Error(res, err.Error(), http.StatusBadRequest)
			} else {
				restutils.Error(res, err.Error(), http.StatusForbidden)
			}
			return
		}
		res.WriteHeader(http.StatusNoContent)
		logger.Info("task deleted")
	}
}

// GetTaskInfo godoc
//
//	@Summary		Get task info
//	@Description	Get task status, date of creation and processing time by id
//	@Tags			Task
//	@Produce		json
//	@Param			id	query		string	true	"task id"
//	@Success		200	{object}	service.TaskInfo
//	@Failure		400	{object}	restutils.HTTPError
//	@Router			/api/task/info [get]
func GetTaskInfo(s taskManager.ServiceInterface) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Info("getting task info")
		id := req.FormValue("id")
		if id == "" {
			logger.Error(ErrorIDRequired)
			restutils.Error(res, ErrorIDRequired.Error(), http.StatusBadRequest)
			return
		}
		taskInfo, err := s.GetTaskInfo(id)
		if err != nil {
			logger.Error(err.Error())
			restutils.Error(res, err.Error(), http.StatusBadRequest)
			return
		}
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(taskInfo)
		logger.Info("task info send")
	}
}
