package apperror

import "errors"

var ErrorIDRequired = errors.New("id required")
var ErrorTaskNotFound = errors.New("task not found")
var ErrorTaskInProcess = errors.New("task in process")
