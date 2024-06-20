package helper

import (
	"e-todo/model/domain"
	"e-todo/model/web"
)

func ToTaskResponse(task domain.Task) web.TaskResponse {
	return web.TaskResponse{
		Id:        task.Id,
		Name:      task.Name,
		Status:    task.Status,
		CreatedAt: task.CreatedAt,
	}
}

func ToTaskResponses(tasks []domain.Task) []web.TaskResponse {
	var taskResponses []web.TaskResponse
	for _,task := range tasks {
		taskResponses = append(taskResponses, ToTaskResponse(task))
	}

	return taskResponses
}
