package service

import (
	"time"
)

type TaskFunction func(args ...interface{})

type IScheduler interface {
	RunTask(name string, function TaskFunction, args ...interface{}) error
	ScheduleRecurrentTask(name string, intervalMs int, ignoreFirstRun bool, function TaskFunction, args ...interface{})
}

type Scheduler struct {
	logger Logger
	IScheduler
}

func (s *Scheduler) RunTask(name string, function TaskFunction, args ...interface{}) {
	function(args...)
}

func (s *Scheduler) ScheduleRecurrentTask(name string, intervalMs int, ignoreFirstRun bool, function TaskFunction, args ...interface{}) {
	t := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
	if !ignoreFirstRun {
		go s.RunTask(name, function, args...)
	}
	for tickerTime := range t.C {
		s.logger.Info("[scheduler/" + name + "] Task started in: " + tickerTime.Format("2006-01-02 15:04:05"))
		go s.RunTask(name, function, args...)

	}
}

func InitScheduler(l Logger) *Scheduler {
	s := Scheduler{logger: l}
	return &s
}
