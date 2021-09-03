package scheduler

import (
	"errors"
	"fmt"
	"testing"
	"tickers-parser/internal/service/logger"
	"tickers-parser/internal/types"
)

func getService() *Scheduler {
	log := logger.NewLogger()
	s := InitScheduler(log)
	return s
}

func TestScheduler_RunTask(t *testing.T) {
	// Succeed
	s := getService()
	testTask := func(args ...interface{}) (interface{}, error) {
		res := fmt.Sprintf("%v %d", args[0], args[1])
		return res, nil
	}
	res, err := s.RunTask("correct_task", testTask, "first", 2)
	if err != nil {
		t.Error(err)
	}
	if !(res == "first 2") {
		t.Error(errors.New("result does not match"))
	}

	res, err = nil, nil

	// Errored
	erroredTask := func(args ...interface{}) (interface{}, error) {
		return nil, errors.New("test error")
	}
	res, err = s.RunTask("incorrect_task", erroredTask)
	if res != nil && err == nil {
		t.Error("an error was not returned")
	}
}

func TestScheduler_ScheduleRecurrentTask(t *testing.T) {
	channels := types.ChannelsPair{
		DataChannel:   make(chan interface{}),
		CancelChannel: make(chan error),
	}
	s := getService()

	// Succeed
	succeedFunc := func(args ...interface{}) (interface{}, error) {
		channels, ok := args[0].(types.ChannelsPair)
		if !ok {
			t.Error(errors.New("wrong argument for channels pair"))
			t.FailNow()
		}
		channels.DataChannel <- "Cool"
		return nil, nil
	}
	go s.ScheduleRecurrentTask("succeed_task", 1*60*1000, false, succeedFunc, channels)

	// Errored
	erroredFunc := func(args ...interface{}) (interface{}, error) {
		channels, ok := args[0].(types.ChannelsPair)
		if !ok {
			t.Error(errors.New("wrong argument for channels pair"))
			t.FailNow()
		}
		err := errors.New("good error")
		channels.CancelChannel <- err
		return nil, err
	}
	go s.ScheduleRecurrentTask("errored_task", 1*60*1000, false, erroredFunc, channels)

	for i := 0; i < 2; i++ {
		select {
		case <-channels.DataChannel:
			s.Logger.Info("data received")
		case <-channels.CancelChannel:
			s.Logger.Info("error received")
		}
	}
}
