package taskManager

import (
	"fmt"
	"sync"
	"time"

	"github.com/beevik/guid"
	"github.com/sater-151/tt-workmate/internal/apperror"
	"github.com/sater-151/tt-workmate/internal/controller/rest/dto"
	client "github.com/sater-151/tt-workmate/internal/pkg"

	logger "github.com/sirupsen/logrus"
)

type TaskManagerService interface {
	CreateTask() string
	StartTask(guid string)
	GetTaskInfo(guid string) (dto.TaskInfo, error)
	DeleteTask(guid string) error
}

type TaskStatus string

const (
	StatusCreated    TaskStatus = "created"
	StatusInProgress TaskStatus = "in progress"
	StatusDone       TaskStatus = "done"
)

type TaskManager struct {
	taskList map[string]*Task
	client   client.TaskManagerPKG
}

type Task struct {
	mu          *sync.RWMutex
	status      TaskStatus
	createDate  time.Time
	processTime string
}

func New() *TaskManager {
	return &TaskManager{
		taskList: make(map[string]*Task),
		client:   client.New(),
	}
}

func (tm *TaskManager) CreateTask() string {
	id := guid.NewString()
	tm.taskList[id] = &Task{
		mu:         new(sync.RWMutex),
		status:     StatusCreated,
		createDate: time.Now(),
	}
	return id
}

func (tm *TaskManager) StartTask(guid string) {
	logger.Info(fmt.Sprintf("task: %s in progress", guid))
	var taskInfo dto.TaskInfo

	tm.taskList[guid].mu.Lock()
	tm.taskList[guid].status = StatusInProgress
	taskInfo.Status = string(tm.taskList[guid].status)
	taskInfo.CreateDate = tm.taskList[guid].createDate.Format(time.ANSIC)
	tm.taskList[guid].mu.Unlock()

	tm.client.SendTask(taskInfo)

	tm.taskList[guid].mu.Lock()
	tm.taskList[guid].processTime = time.Since(tm.taskList[guid].createDate).Round(time.Second).String()
	tm.taskList[guid].status = StatusDone
	tm.taskList[guid].mu.Unlock()
	logger.Info(fmt.Sprintf("task: %s completed", guid))
}

func (tm *TaskManager) GetTaskInfo(guid string) (dto.TaskInfo, error) {
	var taskInfo dto.TaskInfo
	if _, ok := tm.taskList[guid]; !ok {
		return taskInfo, apperror.ErrorTaskNotFound
	}

	tm.taskList[guid].mu.RLock()
	taskInfo.Status = string(tm.taskList[guid].status)
	taskInfo.CreateDate = tm.taskList[guid].createDate.Format(time.ANSIC)
	tm.taskList[guid].mu.RUnlock()

	if taskInfo.Status != string(StatusDone) {
		tm.taskList[guid].mu.RLock()
		taskInfo.ProcessTime = time.Since(tm.taskList[guid].createDate).Round(time.Second).String()
		tm.taskList[guid].mu.RUnlock()
	} else {
		taskInfo.ProcessTime = tm.taskList[guid].processTime
	}
	return taskInfo, nil
}

func (tm *TaskManager) DeleteTask(guid string) error {
	if _, ok := tm.taskList[guid]; !ok {
		return apperror.ErrorTaskNotFound
	}
	if tm.taskList[guid].status != StatusDone {
		return apperror.ErrorTaskInProcess
	}
	delete(tm.taskList, guid)
	return nil
}
