package helper

import (
	"e-todo/model/domain"
	"e-todo/model/web"
)

func ToTaskResponse(task domain.Task) web.TaskResponse {
	return web.TaskResponse{
		Id:          task.Id,
		Name:        task.Name,
		Status:      task.Status,
		Description: task.Description,
		Code:        task.Code,
		CreatedAt:   task.CreatedAt,
	}
}

func ToTaskResponses(tasks []domain.Task) []web.TaskResponse {
	var taskResponses []web.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, ToTaskResponse(task))
	}

	return taskResponses
}

func ToTimerResponse(timer domain.Timer) web.TimerResponse {
	return web.TimerResponse{
		Id:        timer.Id,
		TaskId:    timer.TaskId,
		Time:      timer.Time,
		Title:     timer.Title,
		Status:    timer.Status,
		CreatedAt: timer.CreatedAt,
	}
}

func ToRelationTimerHistoriesResponse(timerHistories domain.RelationTimerHistories) web.RelationTimerHistoriesResponse {
	return web.RelationTimerHistoriesResponse{
		Id:        timerHistories.Id,
		TaskId:    timerHistories.TaskId,
		Time:      timerHistories.Time,
		Status:    timerHistories.Status,
		Histories: timerHistories.Histories,
	}
}

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:        user.Id,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}
}

func ToUserLoginResponse(user domain.User) web.UserLoginResponse {
	return web.UserLoginResponse{
		Id:           user.Id,
		Email:        user.Email,
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
}
