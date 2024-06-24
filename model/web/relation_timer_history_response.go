package web

import "e-todo/model/domain"

type RelationTimerHistoriesResponse struct {
	Id        int                   `json:"id"`
	TaskId    int                   `json:"task_id"`
	Time      string                `json:"time"`
	Status    string                `json:"status"`
	Histories []domain.TimerHistory `json:"posts,omitempty"`
}
