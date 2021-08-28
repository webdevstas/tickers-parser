package scheduler

import (
	"github.com/spf13/viper"
	"runtime"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/service/logger"
	"tickers-parser/internal/service/storage"
	"tickers-parser/internal/service/updater/exchange"
	"tickers-parser/internal/types"
)

type ITasks interface {
	RunTasks()
}

type Tasks struct {
	scheduler *Scheduler
	log       logger.Logger
	storage   *storage.Storage
	config    *viper.Viper
	ITasks
}

func (t *Tasks) RunTasks() {
	t.scheduler.ScheduleRecurrentTask("tickers", t.config.GetInt("app.tickersInterval")*60*1000, true, t.startTickersParsing)
}

func (t *Tasks) startTickersParsing(args ...interface{}) {
	exchanges := exchange.GetExchangesForTickersUpdate()
	tickersChannels := types.ChannelsPair{
		DataChannel:   make(chan interface{}, exchange.ExchangesCount),
		CancelChannel: make(chan error, exchange.ExchangesCount),
	}
	saveChannels := types.ChannelsPair{
		CancelChannel: make(chan error, exchange.ExchangesCount),
		DataChannel:   make(chan interface{}, exchange.ExchangesCount),
	}

	for _, ex := range exchanges {
		go ex.FetchTickers(&tickersChannels)
	}

	for i := 0; i <= exchange.ExchangesCount; i++ {
		select {
		case err := <-tickersChannels.CancelChannel:
			t.log.Error(err)
		case result := <-tickersChannels.DataChannel:
			tickers := result.(entities.ExchangeTickers)
			go t.storage.Save(tickers.Exchange, tickers.Timestamp, tickers.Tickers, saveChannels)
			runtime.Gosched()
			select {
			case err := <-saveChannels.CancelChannel:
				t.log.Error(err)
			case <-saveChannels.DataChannel:
				t.log.Info("[scheduler/tickers] Tickers saved for " + tickers.Exchange)
			}
		}
	}
	close(tickersChannels.DataChannel)
	close(tickersChannels.CancelChannel)
	close(saveChannels.CancelChannel)
	close(saveChannels.DataChannel)
}

func NewTasksService(l logger.Logger, st *storage.Storage, c *viper.Viper) *Tasks {
	t := Tasks{
		scheduler: InitScheduler(l),
		log:       l,
		storage:   st,
		config:    c,
	}
	return &t
}
