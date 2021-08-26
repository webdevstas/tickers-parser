package storage

import (
	"encoding/json"
	"os"
	"strconv"
)

type FileSaver struct {
	rootPath string
}

func (fs *FileSaver) Save(name string, timestamp int64, data interface{}) error {
	err := os.Chdir(fs.rootPath + name)
	if err != nil {
		err = nil
		err = os.Chdir(fs.rootPath)
		if err != nil {
			return err
		}
		err = nil
		err = os.Mkdir(name, 0777)
		err = nil
		err = os.Chdir(fs.rootPath + name)
		if err != nil {
			panic(err)
		}
	}
	file, err := os.Create(strconv.FormatInt(timestamp, 10))
	defer file.Close()
	jsonData, err := json.Marshal(data)
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func NewFileSaver(rootPath string) *FileSaver {
	return &FileSaver{
		rootPath: rootPath,
	}
}
