package service

import (
	"context"
	"e-todo/model/web"
)

type TimerHistoryService interface {
	FindByParentId(ctx context.Context, taskId int) web.RelationTimerHistoriesResponse
	GetWeeklyReport(ctx context.Context) []web.WeeklyReportResponse
}
