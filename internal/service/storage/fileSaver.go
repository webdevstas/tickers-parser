package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"tickers-parser/internal/types"
)

type FileSaver struct {
	workDir  string
	rootPath string
}

func (fs *FileSaver) Save(name string, timestamp int64, data interface{}, channels types.ChannelsPair) {
	cancelChan := channels.CancelChannel
	dataChan := channels.DataChannel
	dataRoot := filepath.FromSlash(fs.workDir + "/" + fs.rootPath + "/")
	err := os.Chdir(dataRoot + name)
	if err != nil {
		err = nil
		err = os.Chdir(dataRoot)
		if err != nil {
			err = os.Mkdir(fs.rootPath, 0777)
			err = os.Chdir(dataRoot)
		}
		err = nil
		err = os.Mkdir(name, 0777)
		if err != nil {
			cancelChan <- err
			return
		}
		err = nil
		err = os.Chdir(dataRoot + name)
		if err != nil {
			cancelChan <- err
			return
		}
	}
	file, err := os.Create(strconv.FormatInt(timestamp, 10) + ".json")
	if err != nil {
		cancelChan <- err
	}
	defer file.Close()
	jsonData, err := json.Marshal(data)
	if err != nil {
		file.Close()
		cancelChan <- err
	}
	_, err = file.Write(jsonData)
	if err != nil {
		file.Close()
		cancelChan <- err
	}
	file.Close()
	dataChan <- true
}

func NewFileSaver(rootPath string) *FileSaver {
	wd, _ := os.Getwd()
	return &FileSaver{
		workDir:  wd,
		rootPath: rootPath,
	}
}
