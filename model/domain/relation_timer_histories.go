package domain

type RelationTimerHistories struct {
	Id        int
	TaskId    int
	Time      string
	Status    string
	Histories []TimerHistory
}
