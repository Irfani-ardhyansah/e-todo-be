package service

import (
	"context"
	"e-todo/model/web"
)

type TimerHistoryService interface {
	FindByParentId(ctx context.Context, timerId int) web.RelationTimerHistoriesResponse
	GetWeeklyReport(ctx context.Context) []web.WeeklyReportResponse
	GetWeeklyReportByTaskId(ctx context.Context, taskId int) []web.WeeklyReportResponse
}
