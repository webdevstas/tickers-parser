package http_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	Client *http.Client
}

func (h *HttpClient) FetchJson(url string, link any) error {
	resp, err := h.Client.Get(url)

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

func GetHttpClient() *HttpClient {
	return &HttpClient{
		Client: &http.Client{},
	}
}
