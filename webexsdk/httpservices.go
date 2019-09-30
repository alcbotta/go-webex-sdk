package webexsdk

import (
	"encoding/json"
	// "fmt"
	// "io/ioutil"
	"net/http"
)

func addBasicHttpHeaders(tokenID string, r *http.Request) {
	r.Header.Set("Authorization", "Bearer "+tokenID)
}

func Get(tokenID, url string, toDecodeInterface interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	addBasicHttpHeaders(tokenID, req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(toDecodeInterface)
	if err != nil {
		return err
	}
	return nil
}
