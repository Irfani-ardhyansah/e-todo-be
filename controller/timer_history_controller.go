package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TimerHistoryController interface {
	ListByTimer(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ListWeeklyReport(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ListWeeklyReportByTaskId(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
