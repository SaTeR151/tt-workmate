package taskManager

import (
	"fmt"
	"sync"
	"time"

	"github.com/beevik/guid"
	"github.com/sater-151/tt-workmate/internal/apperror"
	logger "github.com/sirupsen/logrus"
)

type Service interface {
	CreateTask() string
	StartTask(guid string)
	GetTaskInfo(guid string) (TaskInfo, error)
	DeleteTask(guid string) error
}

var statuses = []string{"created", "in process one", "in process two", "in process three", "in process four", "done"}

type TaskManager struct {
	taskList map[string]*Task
}

type Task struct {
	mu          *sync.RWMutex
	status      string
	createDate  time.Time
	processTime string
}

type TaskInfo struct {
	Status      string `json:"status"`
	CreateDate  string `json:"create_date"`
	ProcessTime string `json:"process_time"`
}

func New() *TaskManager {
	return &TaskManager{
		taskList: make(map[string]*Task),
	}
}

func (tm *TaskManager) CreateTask() string {
	id := guid.NewString()
	tm.taskList[id] = &Task{
		mu:         new(sync.RWMutex),
		status:     statuses[0],
		createDate: time.Now(),
	}
	return id
}

func (tm *TaskManager) StartTask(guid string) {
	logger.Info(fmt.Sprintf("task: %s in process", guid))
	for i := 1; i < 6; i++ {
		time.Sleep(time.Second * 30)
		tm.taskList[guid].mu.Lock()
		tm.taskList[guid].status = statuses[i]
		tm.taskList[guid].mu.Unlock()
	}
	tm.taskList[guid].mu.Lock()
	tm.taskList[guid].processTime = time.Since(tm.taskList[guid].createDate).Round(time.Second).String()
	tm.taskList[guid].mu.Unlock()
	logger.Info(fmt.Sprintf("task: %s completed", guid))
}

func (tm *TaskManager) GetTaskInfo(guid string) (TaskInfo, error) {
	var taskInfo TaskInfo
	if _, ok := tm.taskList[guid]; !ok {
		return taskInfo, apperror.ErrorTaskNotFound
	}

	tm.taskList[guid].mu.RLock()
	taskInfo.Status = tm.taskList[guid].status
	taskInfo.CreateDate = tm.taskList[guid].createDate.Format(time.ANSIC)
	tm.taskList[guid].mu.RUnlock()

	if taskInfo.Status != statuses[len(statuses)-1] {
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
	if tm.taskList[guid].status != statuses[len(statuses)-1] {
		return apperror.ErrorTaskInProcess
	}
	delete(tm.taskList, guid)
	return nil
}
