package impl

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

func sendBytes() error {
	return nil
}


func getBytes(request *http.Request, webClientTimeout int) ([]byte, error) {
	client := &http.Client{Timeout: time.Duration(webClientTimeout) * time.Millisecond}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("provider return error code " + resp.Status)
	}

	return []byte(resp_body), nil
}


func getEtagValue(data []byte) string {
	h := sha1.New()
	h.Write(data)
	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return sha
}


func getJSON(request interface{}) (*string, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	result := string(b)
	return &result, nil
}
