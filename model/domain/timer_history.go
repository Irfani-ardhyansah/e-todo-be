package domain

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
	Date     string
}

type WeeklyReport struct {
	StartDate  string
	EndDate    string
	TotalTime  string
	DataDetail []TaskDetail
}
