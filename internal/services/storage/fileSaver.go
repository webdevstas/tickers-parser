package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
)

type FileSaver struct {
	workDir  string
	rootPath string
}

func (fs *FileSaver) Save(name string, timestamp int64, data interface{}) (bool, error) {
	dataRoot := filepath.FromSlash(fs.workDir + "/" + fs.rootPath + "/")
	err := os.Chdir(dataRoot + name)
	if err != nil {
		err = nil
		err = os.Chdir(dataRoot)
		if err != nil {
			os.Mkdir(fs.rootPath, 0777)
			os.Chdir(dataRoot)
		}
		err = nil
		err = os.Mkdir(name, 0777)
		if err != nil {
			return false, err
		}
		err = nil
		err = os.Chdir(dataRoot + name)
		if err != nil {
			return false, err
		}
	}
	file, err := os.Create(strconv.FormatInt(timestamp, 10) + ".json")
	if err != nil {
		return false, err
	}
	defer file.Close()
	jsonData, err := json.Marshal(data)
	if err != nil {
		file.Close()
		return false, err
	}
	_, err = file.Write(jsonData)
	if err != nil {
		file.Close()
		return false, err
	}
	file.Close()
	return true, nil
}

func NewFileSaver(rootPath string) *FileSaver {
	wd, _ := os.Getwd()
	return &FileSaver{
		workDir:  wd,
		rootPath: rootPath,
	}
}
