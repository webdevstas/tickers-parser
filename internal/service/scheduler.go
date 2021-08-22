package service

import (
	"time"
)

type TaskFunction func(args ...interface{}) error

type IScheduler interface {
	RunTask(name string, function TaskFunction, args ...interface{}) error
	ScheduleRecurrentTask(name string, intervalMs int, ignoreFirstRun bool, function TaskFunction, args ...interface{})
}

type Scheduler struct {
	logger Logger
	IScheduler
}

func (s *Scheduler) RunTask(name string, function TaskFunction, args ...interface{}) error {
	err := function(args...)
	if err != nil {
		s.logger.Error(err)
		return err
	}
	return nil
}

func (s *Scheduler) ScheduleRecurrentTask(name string, intervalMs int, ignoreFirstRun bool, function TaskFunction, args ...interface{}) {
	t := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
	if !ignoreFirstRun {
		err := s.RunTask(name, function, args...)
		if err != nil {
			s.logger.Error(err)
		}
	}
	for tickerTime := range t.C {
		s.logger.Info("[scheduler/" + name + "] Task started in: " + tickerTime.Format("2006-01-02 15:04:05"))
		go func() {
			err := s.RunTask(name, function, args...)
			if err != nil {
				s.logger.Error(err)
			}
		}()
	}
}

func NewScheduler(l Logger) *Scheduler {
	s := Scheduler{logger: l}
	return &s
}
