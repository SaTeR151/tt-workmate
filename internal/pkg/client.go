package client

import (
	"fmt"
	"time"

	"github.com/sater-151/tt-workmate/internal/controller/rest/dto"
)

type TaskManagerPKG interface {
	SendTask(task dto.TaskInfo)
}

type TaskManagerClient struct{}

func New() *TaskManagerClient {
	return &TaskManagerClient{}
}

func (tmc *TaskManagerClient) SendTask(task dto.TaskInfo) {
	fmt.Println(1)
	time.Sleep(time.Minute * 3)
}
