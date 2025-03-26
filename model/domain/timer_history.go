package domain

import "time"

type TimerHistory struct {
	Id        int
	TimerId   int
	Status    string
	TimeLog   string
	CreatedAt string
}

type RelationTimerHistories struct {
	Id        int
	TaskId    int
	Time      string
	Status    string
	Histories []TimerHistory
}

type TaskDetail struct {
	Id       int
	TaskName string
	Time     string
	Date     time.Time
}

type GroupedByDate struct {
	Date         string
	DataGroupded []TaskSummary
}

type TaskSummary struct {
	Id       int
	TaskName string
	Time     string
}
