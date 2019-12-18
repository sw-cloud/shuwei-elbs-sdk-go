package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"shuwei-elbs-sdk-go/log"
)

func HttpPost(url string, reqBody []byte, authorization string) ([]byte, error) {
	buffer := bytes.NewBuffer(reqBody)
	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		log.Errorf("http NewRequest: %v\n", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Authorization", authorization)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("http post: %v\n", err)
		return nil, err
	}
	defer  resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("http read resp: %v\n", err)
		return nil, err
	}
	return respBody, nil
}
