package scheduler

import (
	"errors"
	"fmt"
	"testing"
	"tickers-parser/internal/service/logger"
)

func getService() IScheduler {
	log := logger.NewLogger()
	s := InitScheduler(log)
	return s
}

func TestCorrectRunTask(t *testing.T) {
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
}

func TestIncorrectRunTask(t *testing.T) {
	s := getService()
	testTask := func(args ...interface{}) (interface{}, error) {
		return nil, errors.New("test error")
	}
	res, err := s.RunTask("incorrect_task", testTask)
	if res != nil && err == nil {
		t.Error("error not returns")
	}
}
