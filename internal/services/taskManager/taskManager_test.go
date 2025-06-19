package taskManager

import (
	"testing"
	"time"

	"github.com/sater-151/tt-workmate/internal/controller/rest/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (tmc *MockClient) SendTask(task dto.TaskInfo) {
	time.Sleep(time.Second * 2)
}

func TestStartTask(t *testing.T) {
	clientMock := new(MockClient)
	clientMock.On("SendTask", mock.Anything)
	taskManager := New()
	taskManager.client = clientMock

	taskId := taskManager.CreateTask()

	assert.Equal(t, taskManager.taskList[taskId].status, StatusCreated)

	go taskManager.StartTask(taskId)
	time.Sleep(time.Second)
	assert.Equal(t, taskManager.taskList[taskId].status, StatusInProgress)
	time.Sleep(time.Second * 2)

	assert.Equal(t, taskManager.taskList[taskId].status, StatusDone)
}

func TestDeleteTask(t *testing.T) {
	clientMock := new(MockClient)
	clientMock.On("SendTask", mock.Anything)
	taskManager := New()
	taskManager.client = clientMock

	taskId := taskManager.CreateTask()

	go taskManager.StartTask(taskId)
	time.Sleep(time.Second)
	assert.Error(t, taskManager.DeleteTask(taskId))
	assert.Error(t, taskManager.DeleteTask("unknown id"))
	time.Sleep(time.Second * 2)

	assert.NoError(t, taskManager.DeleteTask(taskId))
}

func TestGetTaskInfo(t *testing.T) {
	clientMock := new(MockClient)
	clientMock.On("SendTask", mock.Anything)
	taskManager := New()
	taskManager.client = clientMock
	_, err := taskManager.GetTaskInfo("unknown id")
	assert.Error(t, err)

	taskId := taskManager.CreateTask()

	go taskManager.StartTask(taskId)
	time.Sleep(time.Second)
	task, err := taskManager.GetTaskInfo(taskId)
	assert.NoError(t, err)
	assert.Equal(t, task.Status, string(StatusInProgress))
	time.Sleep(time.Second * 2)

	task, err = taskManager.GetTaskInfo(taskId)
	assert.NoError(t, err)
	assert.Equal(t, task.Status, string(StatusDone))
	assert.NoError(t, taskManager.DeleteTask(taskId))
}
