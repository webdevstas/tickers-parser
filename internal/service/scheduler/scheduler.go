package scheduler

import (
	"tickers-parser/internal/service/logger"
	"time"
)

type TaskFunction func(args ...interface{})

type IScheduler interface {
	RunTask(name string, function TaskFunction, args ...interface{}) error
	ScheduleRecurrentTask(name string, intervalMs int, ignoreFirstRun bool, function TaskFunction, args ...interface{})
}

type Scheduler struct {
	logger logger.Logger
	IScheduler
}

func (s *Scheduler) RunTask(name string, function TaskFunction, args ...interface{}) {
	s.logger.Info("[scheduler/" + name + "] Task started")
	function(args...)
	s.logger.Info("[scheduler/" + name + "] Task ended")
}

func (s *Scheduler) ScheduleRecurrentTask(name string, intervalMs int, ignoreFirstRun bool, function TaskFunction, args ...interface{}) {
	t := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
	if !ignoreFirstRun {
		s.RunTask(name, function, args...)
	}
	for _ = range t.C {
		s.RunTask(name, function, args...)
	}
}

func InitScheduler(l logger.Logger) *Scheduler {
	return &Scheduler{logger: l}
}
