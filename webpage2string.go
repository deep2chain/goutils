package goutils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getWebSite(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("STATUS rror: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("READ body: %v", err)
	}

	return data, nil
}

func GetWebSite(url string) ([]byte, error) {
	return getWebSite(url)
}
