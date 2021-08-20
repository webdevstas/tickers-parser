package scheduler

import (
	"tickers-parser/modules"
)

type taskFunction func(args ...interface{}) (interface{}, error)

type taskParams struct {
	name     string
	function taskFunction
	interval int
	args     interface{}
}

type IScheduler interface {
	runTask(name string, function taskFunction, args ...interface{}) interface{}
}

type Scheduler struct {
	logger modules.Logger
	IScheduler
}

func (s *Scheduler) RunTask(name string, function taskFunction, args ...interface{}) (interface{}, error) {
	s.logger.Info("Run task: " + name)
	res, err := function(args)
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	return res, nil
}

func NewSchedulerModule(l modules.Logger) *Scheduler {
	s := Scheduler{logger: l}
	return &s
}
