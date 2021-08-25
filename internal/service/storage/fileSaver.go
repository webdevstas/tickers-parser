package storage

import "fmt"

type FileSaver struct {
}

func (fs *FileSaver) Save(name string, timestamp int64, data interface{}) error {
	fmt.Print(name, timestamp, data)
	return nil //TODO: Реализовать сохранение в файлы
}

func NewFileSaver(rootPath string) *FileSaver {
	return &FileSaver{}
}
