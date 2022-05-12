package scheduler

import (
	"github.com/spf13/viper"
	"runtime"
	"tickers-parser/internal/entities"
	"tickers-parser/internal/repository"
	"tickers-parser/internal/services/logger"
	"tickers-parser/internal/services/storage"
	"tickers-parser/internal/services/updater"
	"tickers-parser/internal/types"
)

type ITasks interface {
	RunTasks()
}

type Tasks struct {
	scheduler    *Scheduler
	log          logger.Logger
	config       *viper.Viper
	repository   *repository.Repositories
	tickersStore storage.TickersStore
}

func (t *Tasks) RunTasks() {
	t.scheduler.ScheduleRecurrentTask("tickers", t.config.GetInt("app.tickersInterval")*60*1000, false, t.startTickersParsing)
}

func (t *Tasks) startTickersParsing(args ...interface{}) (interface{}, error) {
	exchanges := updater.GetExchangesForTickersUpdate(t.repository)
	exchangesCount := len(exchanges)

	tickersChannels := types.ChannelsPair[entities.ExchangeTickers]{
		DataChannel:   make(chan entities.ExchangeTickers, exchangesCount),
		CancelChannel: make(chan error, exchangesCount),
	}

	saveChannels := types.ChannelsPair[interface{}]{
		CancelChannel: make(chan error, exchangesCount),
		DataChannel:   make(chan interface{}, exchangesCount),
	}

	defer func() {
		tickersChannels.CloseAll()
		saveChannels.CloseAll()
		t.log.Info("[scheduler/tickers] All channels closed")
	}()

	for _, ex := range exchanges {
		go func(exchange entities.Exchange, channels types.ChannelsPair[entities.ExchangeTickers]) {
			exchangeTickers := entities.ExchangeTickers{
				Exchange: exchange,
			}
			res, err := exchange.FetchTickers()
			if err != nil {
				channels.CancelChannel <- err
				return
			}
			exchangeTickers.Tickers = res
			channels.DataChannel <- exchangeTickers
		}(ex, tickersChannels)
	}

	for i := 0; i < exchangesCount; i++ {
		select {
		case err := <-tickersChannels.CancelChannel:
			t.log.Warn(err)
			continue
		case result := <-tickersChannels.DataChannel:
			tickers := result
			go func(channels types.ChannelsPair[interface{}]) {
				res, err := t.tickersStore.SaveTickersForExchange(tickers.Exchange.ID, tickers.Tickers)
				if err != nil {
					channels.CancelChannel <- err
				} else {
					channels.DataChannel <- res
				}
			}(saveChannels)
			select {
			case err := <-saveChannels.CancelChannel:
				t.log.Warn(err)
				continue
			case <-saveChannels.DataChannel:
				t.log.Info("[scheduler/tickers] Tickers saved for " + tickers.Exchange.Key)
				runtime.Gosched()
			}
		}
	}
	return nil, nil
}

func NewTasksService(l logger.Logger, c *viper.Viper, r *repository.Repositories) *Tasks {
	t := Tasks{
		scheduler:    InitScheduler(l),
		log:          l,
		config:       c,
		repository:   r,
		tickersStore: storage.NewTickersStoreService(r),
	}
	return &t
}
