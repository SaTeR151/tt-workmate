package rest

import (
	"encoding/json"
	"net/http"

	"github.com/sater-151/tt-workmate/internal/apperror"
	"github.com/sater-151/tt-workmate/internal/controller/rest/dto"
	"github.com/sater-151/tt-workmate/internal/controller/rest/restutils"
	"github.com/sater-151/tt-workmate/internal/services/taskManager"
	logger "github.com/sirupsen/logrus"
)

// CreateTask godoc
//
//	@Summary		Create a new task
//	@Tags			Task
//	@Produce		json
//	@Success		201	{object}	dto.ResponceId
//	@Router			/api/task/new [post]
func CreateTask(s taskManager.TaskManagerService) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Info("creating task")
		id := s.CreateTask()
		logger.Info("task created")
		go s.StartTask(id)
		res.WriteHeader(http.StatusCreated)
		json.NewEncoder(res).Encode(dto.ResponceId{Id: id})
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
//	@Success		204	{object}	dto.ResponceId
//	@Failure		400	{object}	restutils.HTTPError
//	@Failure		403	{object}	restutils.HTTPError
//	@Router			/api/task/delete [delete]
func DeleteTask(s taskManager.TaskManagerService) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Info("deleting task")
		id := req.FormValue("id")
		if id == "" {
			logger.Error(apperror.ErrorIDRequired)
			restutils.Error(res, apperror.ErrorIDRequired.Error(), http.StatusBadRequest)
			return
		}
		if err := s.DeleteTask(id); err != nil {
			logger.Warn(err.Error())
			if err == apperror.ErrorTaskNotFound {
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
//	@Success		200	{object}	taskManager.TaskInfo
//	@Failure		400	{object}	restutils.HTTPError
//	@Router			/api/task/info [get]
func GetTaskInfo(s taskManager.TaskManagerService) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		logger.Info("getting task info")
		id := req.FormValue("id")
		if id == "" {
			logger.Error(apperror.ErrorIDRequired)
			restutils.Error(res, apperror.ErrorIDRequired.Error(), http.StatusBadRequest)
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
