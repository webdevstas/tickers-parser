package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchJson[T any](url string, link T) error {
	resp, err := http.Get(url)

	if err != nil || resp.StatusCode != 200 {
		err = fmt.Errorf("error to get %v. Code %d. Error: %v", url, resp.StatusCode, err)
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		err = fmt.Errorf("error to parse body %v. %v", url, err)
		return err
	}

	err = json.Unmarshal(body, &link)

	if err != nil {
		err = fmt.Errorf("unmarshal error %v. %v", url, err)
		return err
	}

	return nil
}
