package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
	"tickers-parser/internal/types"
)

type FileSaver struct {
	rootPath string
}

func (fs *FileSaver) Save(name string, timestamp int64, data interface{}, channels types.ChannelsPair) {
	cancelChan := channels.CancelChannel
	wd, err := os.Getwd()
	if err != nil {
		cancelChan <- err
	}
	err = nil
	dataRoot := filepath.FromSlash(wd + "/" + fs.rootPath + "/")
	err = os.Chdir(dataRoot + name)
	if err != nil {
		err = nil
		err = os.Chdir(dataRoot)
		if err != nil {
			cancelChan <- err
		}
		err = nil
		err = os.Mkdir(name, 0777)
		if err != nil {
			cancelChan <- err
		}
		err = nil
		err = os.Chdir(dataRoot + name)
		if err != nil {
			cancelChan <- err
		}
	}
	file, err := os.Create(strconv.FormatInt(timestamp, 10) + ".json")
	if err != nil {
		cancelChan <- err
	}
	defer file.Close()
	jsonData, err := json.Marshal(data)
	if err != nil {
		cancelChan <- err
	}
	_, err = file.Write(jsonData)
	if err != nil {
		cancelChan <- err
	}
	file.Close()
	os.Chdir("../../")
}

func NewFileSaver(rootPath string) *FileSaver {
	return &FileSaver{
		rootPath: rootPath,
	}
}
