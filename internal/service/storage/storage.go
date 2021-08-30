package storage

type IStorage interface {
	Save(name string, timestamp int64, data interface{}) (bool, error)
}

type Storage struct {
	SaveService IStorage
}

func (s *Storage) Save(name string, timestamp int64, data interface{}) (bool, error) {
	res, err := s.SaveService.Save(name, timestamp, data)
	if err != nil {
		return res, err
	}
	return res, nil
}

func NewStorageService(saveService IStorage) *Storage {
	return &Storage{
		SaveService: saveService,
	}
}
