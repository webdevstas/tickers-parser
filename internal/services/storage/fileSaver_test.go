package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

type testData struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func TestFileSaver_Save(t *testing.T) {
	wd, _ := os.Getwd()
	rootPath := "test_data"
	dirname := "testing"
	timestamp := time.Now().Unix()
	// save
	s := NewFileSaver(rootPath)
	data := testData{
		Name: "Stas",
		Age:  32,
	}
	_, err := s.Save(dirname, timestamp, data)
	if err != nil {
		t.Error(err)
	}
	// check
	filePath := fmt.Sprintf("%v/%v/%v/%v.json", wd, rootPath, dirname, timestamp)
	filePath = filepath.FromSlash(filePath)
	file, err := os.ReadFile(filePath)
	if err != nil {
		t.Error(err)
	}
	if bytesData, _ := json.Marshal(data); bytes.Compare(file, bytesData) != 0 {
		t.Error("data not eq")
	}
	//clean up
	err = os.RemoveAll(fmt.Sprintf("%v/%v", wd, rootPath))
	if err != nil {
		t.Error(err)
	}
}
