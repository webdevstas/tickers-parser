package scheduler

import (
	"errors"
	"fmt"
	"testing"
	"tickers-parser/internal/service/logger"
)

func TestCorrectRunTask(t *testing.T) {
	log := logger.NewLogger()
	s := InitScheduler(log)
	testTask := func(args ...interface{}) (interface{}, error) {
		res := fmt.Sprintf("%v %d", args[0], args[1])
		return res, nil
	}
	res, err := s.RunTask("test", testTask, "first", 2)
	if err != nil {
		t.Error(err)
	}
	if !(res == "first 2") {
		t.Error(errors.New("result does not match"))
	}
}
