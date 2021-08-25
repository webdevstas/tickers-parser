package storage

type IStorage interface {
	Save(name string, timestamp int64, data interface{}) error
}

type Storage struct {
	SaveService IStorage
}

func (s *Storage) Save(name string, timestamp int64, data interface{}) error {
	err := s.SaveService.Save(name, timestamp, data)
	if err != nil {
		return err
	}
	return nil
}

func NewStorageService(saveService IStorage) *Storage {
	return &Storage{
		SaveService: saveService,
	}
}
