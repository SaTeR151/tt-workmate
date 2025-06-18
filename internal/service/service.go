package service

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/beevik/guid"
	logger "github.com/sirupsen/logrus"
)

type ServiceInterface interface {
	CreateTask() string
	Processing(guid string)
	GetTaskInfo(guid string) (TaskInfo, error)
	DeleteTask(guid string) error
}

var ErrorTaskNotFound = errors.New("task not found")
var ErrorTaskInProcess = errors.New("task in process")
var statuses = []string{"created", "in process one", "in process two", "in process three", "in process four", "done"}

type Service struct {
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

func CreateTaskList() *Service {
	return &Service{
		taskList: make(map[string]*Task),
	}
}

func (s *Service) CreateTask() string {
	id := guid.NewString()
	s.taskList[id] = &Task{
		mu:         new(sync.RWMutex),
		status:     statuses[0],
		createDate: time.Now(),
	}
	return id
}

func (s *Service) Processing(guid string) {
	logger.Info(fmt.Sprintf("task: %s in process", guid))
	for i := 1; i < 6; i++ {
		time.Sleep(time.Second * 5)
		s.taskList[guid].mu.Lock()
		s.taskList[guid].status = statuses[i]
		s.taskList[guid].mu.Unlock()
	}
	s.taskList[guid].mu.Lock()
	s.taskList[guid].processTime = time.Since(s.taskList[guid].createDate).Round(time.Second).String()
	s.taskList[guid].mu.Unlock()
	logger.Info(fmt.Sprintf("task: %s completed", guid))
}

func (s *Service) GetTaskInfo(guid string) (TaskInfo, error) {
	var taskInfo TaskInfo
	if _, ok := s.taskList[guid]; !ok {
		return taskInfo, ErrorTaskNotFound
	}

	s.taskList[guid].mu.RLock()
	taskInfo.Status = s.taskList[guid].status
	taskInfo.CreateDate = s.taskList[guid].createDate.Format(time.ANSIC)
	s.taskList[guid].mu.RUnlock()

	if taskInfo.Status != statuses[len(statuses)-1] {
		s.taskList[guid].mu.RLock()
		taskInfo.ProcessTime = time.Since(s.taskList[guid].createDate).Round(time.Second).String()
		s.taskList[guid].mu.RUnlock()
	} else {
		taskInfo.ProcessTime = s.taskList[guid].processTime
	}
	return taskInfo, nil
}

func (s *Service) DeleteTask(guid string) error {
	if _, ok := s.taskList[guid]; !ok {
		return ErrorTaskNotFound
	}
	if s.taskList[guid].status != statuses[len(statuses)-1] {
		return ErrorTaskInProcess
	}
	delete(s.taskList, guid)
	return nil
}
