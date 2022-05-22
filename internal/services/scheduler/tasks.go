package scheduler

import (
	"context"
	"runtime"
	"sync"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/services/logger"
	"tickers-parser/internal/services/updater"

	"github.com/spf13/viper"
)

type ITasks interface {
	RunTasks()
}

type Tasks struct {
	scheduler    *Scheduler
	log          logger.Logger
	config       *viper.Viper
	repository   repository.IRepository
	tickersStore updater.TickersStore
}

func NewTasksService(l logger.Logger, c *viper.Viper, r *repository.Repository) *Tasks {
	t := Tasks{
		scheduler:    InitScheduler(l),
		log:          l,
		config:       c,
		repository:   r,
		tickersStore: updater.NewTickersStoreService(r),
	}
	return &t
}

func (t *Tasks) RunTasks() {
	t.scheduler.ScheduleRecurrentTask("tickers", t.config.GetInt("app.tickersInterval")*60*1000, false, t.startTickersParsing)
}

func (t *Tasks) startTickersParsing(args ...interface{}) (interface{}, error) {
	exchanges := t.repository.GetExchangesForTickersUpdate()
	wg := &sync.WaitGroup{}
	ctx := context.Background()
	var routinesNumber int

	if len(exchanges) < runtime.NumCPU() {
		routinesNumber = len(exchanges)
	} else {
		routinesNumber = runtime.NumCPU()
	}

	inpChan := make(chan entities.Exchange, routinesNumber)
	outChan := make(chan entities.ExchangeTickers, routinesNumber)
	errChan := make(chan error, routinesNumber)

	for i := 0; i <= routinesNumber; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, inpChan, outChan, errChan)
		}()
	}

	go func() {
		for _, ex := range exchanges {
			inpChan <- ex
		}
		close(inpChan)
	}()

	go func() {
		wg.Wait()
		close(outChan)
		close(errChan)
	}()

	for {
		select {
		case res, ok := <-outChan:
			if !ok {
				return nil, nil
			}
			go func(tickersResult entities.ExchangeTickers) {
				ok, err := t.repository.SaveTickersForExchange(tickersResult.Exchange.ID, tickersResult.Tickers)
				if !ok {
					t.log.Error(err)
					return
				}
				t.log.Info("saved tickers for ", res.Exchange.Name)
			}(res)
		case err := <-errChan:
			if err != nil {
				t.log.Error(err)
			}
		}
	}
}

func worker(ctx context.Context, inpChan chan entities.Exchange, outChan chan entities.ExchangeTickers, errChan chan error) {
	for {
		select {
		case <-ctx.Done():
			return
		case ex, ok := <-inpChan:
			if !ok {
				return
			}
			tickers, err := ex.FetchTickers()
			if err != nil {
				errChan <- err
				return
			}
			outChan <- entities.ExchangeTickers{
				Exchange: ex,
				Tickers:  tickers,
			}
		}
	}
}
