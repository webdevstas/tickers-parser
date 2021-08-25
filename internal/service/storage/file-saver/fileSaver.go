package file_saver

import (
	"encoding/json"
	"os"
	"strconv"
	"tickers-parser/internal/service"
)

type FileSaver struct {
	rootPath string
	log      service.Logger
}

func (fs *FileSaver) Save(name string, timestamp int64, data interface{}) error {
	err := os.Chdir(fs.rootPath + name)
	if err != nil {
		fs.log.Warn(err)
		err = os.Mkdir(fs.rootPath+name, os.ModeDir)
		if err != nil {
			fs.log.Error(err)
		}
		err = os.Chdir(fs.rootPath + name)
		if err != nil {
			fs.log.Error(err)
			panic(err)
		}
	}
	file := os.NewFile(uintptr(timestamp), strconv.FormatInt(timestamp, 10))
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
