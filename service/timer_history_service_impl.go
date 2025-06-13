package service

import (
	"context"
	"database/sql"
	exception "e-todo/excception"
	"e-todo/helper"
	"e-todo/model/domain"
	"e-todo/model/web"
	"e-todo/repository"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type TimerHistoryServiceImpl struct {
	TimerHistoryRepository repository.TimerHistoryRepository
	DB                     *sql.DB
	Validate               *validator.Validate
}

func NewTimerHistoryService(timerHistoryRepository repository.TimerHistoryRepository, DB *sql.DB, validate *validator.Validate) TimerHistoryService {
	return &TimerHistoryServiceImpl{
		TimerHistoryRepository: timerHistoryRepository,
		DB:                     DB,
		Validate:               validate,
	}
}

func (service *TimerHistoryServiceImpl) FindByParentId(ctx context.Context, timerId int) web.RelationTimerHistoriesResponse {
	timer, err := service.TimerHistoryRepository.FindByParentId(ctx, service.DB, timerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToRelationTimerHistoriesResponse(timer)
}

func (service *TimerHistoryServiceImpl) GetWeeklyReport(ctx context.Context) []web.WeeklyReportResponse {
	tasks, err := service.TimerHistoryRepository.GetAll(ctx, service.DB)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	weeklyData := make(map[string][]domain.TaskDetail)

	for _, task := range tasks {
		startDate, endDate := getWeekRange(task.Date)
		weekKey := fmt.Sprintf("%s - %s", startDate.Format("02-01-2006"), endDate.Format("02-01-2006"))
		weeklyData[weekKey] = append(weeklyData[weekKey], task)
	}

	var reports []web.WeeklyReportResponse

	for weekKey, tasks := range weeklyData {
		var startDate, endDate string
		fmt.Sscanf(weekKey, "%s - %s", &startDate, &endDate)

		totalTime := calculateTotalTime(tasks)

		groupedTaskByDate := groupTaskByDate(tasks)

		reports = append(reports, web.WeeklyReportResponse{
			StartDate:  startDate,
			EndDate:    endDate,
			TotalTime:  totalTime,
			DataDetail: groupedTaskByDate,
		})
	}

	return reports
}

func (service *TimerHistoryServiceImpl) GetWeeklyReportByTaskId(ctx context.Context, taskId int) []web.WeeklyReportResponse {
	tasks, err := service.TimerHistoryRepository.GetByTaskId(ctx, service.DB, taskId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	weeklyData := make(map[string][]domain.TaskDetail)

	for _, task := range tasks {
		startDate, endDate := getWeekRange(task.Date)
		weekKey := fmt.Sprintf("%s - %s", startDate.Format("02-01-2006"), endDate.Format("02-01-2006"))
		weeklyData[weekKey] = append(weeklyData[weekKey], task)
	}

	var reports []web.WeeklyReportResponse

	for weekKey, tasks := range weeklyData {
		var startDate, endDate string
		fmt.Sscanf(weekKey, "%s - %s", &startDate, &endDate)

		totalTime := calculateTotalTime(tasks)

		groupedTaskByDate := groupTaskByDate(tasks)

		reports = append(reports, web.WeeklyReportResponse{
			StartDate:  startDate,
			EndDate:    endDate,
			TotalTime:  totalTime,
			DataDetail: groupedTaskByDate,
		})
	}

	return reports
}

func groupTaskByDate(tasks []domain.TaskDetail) []web.GroupedByDateResponse {
	grouped := make(map[string][]web.TaskSummaryResponse)

	for _, task := range tasks {
		dateKey := task.Date.Format("02-01-2006")

		grouped[dateKey] = append(grouped[dateKey], web.TaskSummaryResponse{
			Id:       task.Id,
			TaskName: task.TaskName,
			Time:     task.Time,
		})
	}

	results := []web.GroupedByDateResponse{}
	for date, tasks := range grouped {
		results = append(results, web.GroupedByDateResponse{
			Date:        date,
			DataGrouped: tasks,
		})
	}

	return results
}

func getWeekRange(date time.Time) (time.Time, time.Time) {
	weekday := date.Weekday()
	startDate := date.AddDate(0, 0, -int(weekday)+1)
	endDate := startDate.AddDate(0, 0, 6)

	return startDate, endDate
}

func calculateTotalTime(tasks []domain.TaskDetail) string {
	var totalDuration time.Duration

	for _, task := range tasks {
		duration, err := parseTimeToDuration(task.Time)
		if err != nil {
			fmt.Println("Error parsing duration:", err)
			continue
		}
		totalDuration += duration
	}

	hours := int(totalDuration.Hours())
	minutes := int(totalDuration.Minutes()) % 60
	seconds := int(totalDuration.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func parseTimeToDuration(timeStr string) (time.Duration, error) {
	var hours, minutes, seconds int
	_, err := fmt.Sscanf(timeStr, "%d:%d:%d", &hours, &minutes, &seconds)
	if err != nil {
		return 0, err
	}

	return time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second, nil
}
