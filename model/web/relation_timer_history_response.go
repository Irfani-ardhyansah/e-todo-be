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
	StartDate  string                  `json:"start_date"`
	EndDate    string                  `json:"end_date"`
	TotalTime  string                  `json:"total_time"`
	DataDetail []GroupedByDateResponse `json:"data_detail"`
}

type GroupedByDateResponse struct {
	Date        string                `json:"date"`
	DataGrouped []TaskSummaryResponse `json:"data_grouped"`
}
type TaskSummaryResponse struct {
	Id       int    `json:"id"`
	TaskName string `json:"task_name"`
	Time     string `json:"time"`
}
