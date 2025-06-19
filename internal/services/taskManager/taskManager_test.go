package taskManager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartTask(t *testing.T) {
	taskManager := New()
	taskId := taskManager.CreateTask()
	assert.Equal(taskManager.taskList[taskId].status)
}
