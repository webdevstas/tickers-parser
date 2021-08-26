package storage

import "tickers-parser/internal/types"

type IStorage interface {
	Save(name string, timestamp int64, data interface{}, channels types.ChannelsPair)
}

type Storage struct {
	SaveService IStorage
}

func (s *Storage) Save(name string, timestamp int64, data interface{}, channels types.ChannelsPair) {
	s.SaveService.Save(name, timestamp, data, channels)
}

func NewStorageService(saveService IStorage) *Storage {
	return &Storage{
		SaveService: saveService,
	}
}
