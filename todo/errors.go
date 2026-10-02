package todo

import "errors"

var TaskNotFound = errors.New("task not found")
var ErrTaskAlreadyExists = errors.New("task already exists")
