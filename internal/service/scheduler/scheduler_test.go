package scheduler

import (
	"errors"
	"fmt"
	"sync"
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

func TestScheduler_ScheduleRecurrentTaskSucceed(t *testing.T) {
	channels := types.ChannelsPair{
		DataChannel:   make(chan interface{}),
		CancelChannel: make(chan error),
	}
	s := getService()
	var wg sync.WaitGroup
	wg.Add(1)
	succeedFunc := func(args ...interface{}) (interface{}, error) {
		channels, ok := args[0].(types.ChannelsPair)
		wg := args[1].(*sync.WaitGroup)
		if !ok {
			t.Error(errors.New("wrong argument for channels pair"))
		}
		wg.Done()
		channels.DataChannel <- "Cool"
		return nil, nil
	}
	go s.ScheduleRecurrentTask("succeed_task", 1*60*1000, false, succeedFunc, channels, &wg)
	wg.Wait()
	select {
	case <-channels.DataChannel:
		s.Logger.Info("ok")
	default:
		t.Error("no data received from function")
	}
}

func TestScheduler_ScheduleRecurrentTaskErrored(t *testing.T) {
	channels := types.ChannelsPair{
		DataChannel:   make(chan interface{}),
		CancelChannel: make(chan error),
	}
	s := getService()
	var wg sync.WaitGroup
	wg.Add(1)
	succeedFunc := func(args ...interface{}) (interface{}, error) {
		channels, ok := args[0].(types.ChannelsPair)
		wg := args[1].(*sync.WaitGroup)
		if !ok {
			t.Error(errors.New("wrong argument for channels pair"))
		}
		wg.Done()
		err := errors.New("good error")
		channels.CancelChannel <- err
		return nil, err
	}
	go s.ScheduleRecurrentTask("succeed_task", 1*60*1000, false, succeedFunc, channels, &wg)
	wg.Wait()
	select {
	case <-channels.CancelChannel:
		s.Logger.Info("ok")
	default:
		t.Error("no data received from function")
	}
}
