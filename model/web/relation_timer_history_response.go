package web

import "e-todo/model/domain"

type RelationTimerHistoriesResponse struct {
	Id        int                   `json:"id"`
	TaskId    int                   `json:"task_id"`
	Time      string                `json:"time"`
	Status    string                `json:"status"`
	Histories []domain.TimerHistory `json:"posts,omitempty"`
}

type WeeklyReportResponse struct {
	StartDate  string                 `json:"start_date"`
	EndDate    string                 `json:"end_date"`
	TotalTime  string                 `json:"total_time"`
	DataDetail []domain.GroupedByDate `json:"data_detail"`
}
