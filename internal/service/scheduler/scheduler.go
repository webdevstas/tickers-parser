package scheduler

import (
	"strconv"
	"tickers-parser/internal/service/logger"
	"time"
)

type TaskFunction func(args ...interface{}) (interface{}, error)

type IScheduler interface {
	RunTask(name string, function TaskFunction, args ...interface{}) (interface{}, error)
	ScheduleRecurrentTask(name string, intervalMs int, ignoreFirstRun bool, function TaskFunction, args ...interface{})
}

type Scheduler struct {
	logger logger.Logger
	IScheduler
}

func (s *Scheduler) RunTask(name string, function TaskFunction, args ...interface{}) (interface{}, error) {
	s.logger.Info("[scheduler/" + name + "] Task started")
	start := time.Now()
	res, err := function(args...)
	if err != nil {
		return nil, err
	}
	end := time.Since(start).Milliseconds()
	s.logger.Info("[scheduler/" + name + "] Task ended in " + strconv.FormatInt(end, 10) + "ms")
	return res, nil
}

func (s *Scheduler) ScheduleRecurrentTask(name string, intervalMs int, ignoreFirstRun bool, function TaskFunction, args ...interface{}) {
	t := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
	if !ignoreFirstRun {
		go func() {
			_, err := s.RunTask(name, function, args...)
			if err != nil {
				s.logger.Error(err)
			}
		}()
	}
	for _ = range t.C {
		go func() {
			_, err := s.RunTask(name, function, args...)
			if err != nil {
				s.logger.Error(err)
			}
		}()
	}
}

func InitScheduler(l logger.Logger) *Scheduler {
	return &Scheduler{logger: l}
}
