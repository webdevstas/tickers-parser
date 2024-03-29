package scheduler

import (
	"runtime"
	"strconv"
	"tickers-parser/internal/services/logger"
	"time"
)

type TaskFunction func(args ...interface{}) (interface{}, error)

type IScheduler interface {
	RunTask(name string, function TaskFunction, args ...interface{}) (interface{}, error)
	ScheduleRecurrentTask(name string, intervalMs float64, ignoreFirstRun bool, function TaskFunction, args ...interface{})
}

type Scheduler struct {
	Logger logger.ILogger
}

func (s *Scheduler) RunTask(name string, function TaskFunction, args ...interface{}) (interface{}, error) {
	s.Logger.Info("[scheduler/" + name + "] Task started")
	start := time.Now()
	res, err := function(args...)
	if err != nil {
		return nil, err
	}
	end := time.Since(start).Seconds()
	s.Logger.Info("[scheduler/" + name + "] Task ended in " + strconv.FormatFloat(end, 'f', 3, 64) + "s")
	return res, nil
}

func (s *Scheduler) ScheduleRecurrentTask(name string, intervalMs float64, ignoreFirstRun bool, function TaskFunction, args ...interface{}) {
	t := time.NewTicker(time.Duration(intervalMs) * time.Millisecond)
	defer t.Stop()
	if !ignoreFirstRun {
		go func() {
			_, err := s.RunTask(name, function, args...)
			if err != nil {
				s.Logger.Error(err)
			}
		}()
	}
	for range t.C {
		go func() {
			_, err := s.RunTask(name, function, args...)
			if err != nil {
				s.Logger.Error(err)
			}
			runtime.Gosched()
		}()
	}
}

func InitScheduler(l *logger.Logger) *Scheduler {
	return &Scheduler{Logger: l}
}
