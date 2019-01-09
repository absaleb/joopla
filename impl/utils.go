package impl

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
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

func getBytesByMethod(request interface{}, method ZooplaMethod) ([]byte, error){
	if method == Branch_update {
		request = ZooplaBranchUpdateRequest{
			BranchName:      ZooplaBranchName,
			BranchReference: ZooplaBranchReference,
			Email:           ZooplaBranchEmail,
			Location: &Location{
				Coordinates: &Coordinates{Longitude: 0.0, Latitude: 0.0},
				CountryCode: ZooplaBranchCountryCode,
				PostalCode:  ZooplaBranchPostalCode,
				StreetName:  ZooplaBranchStreeName,
				TownOrCity:  ZooplaBranchTownOrCity,
			},
			Telephone: ZooplaBranchTelephone,
			Website:   ZooplaBranchWebsite,
		}
	}
	jsn, err := getJSON(request)
	if err != nil {
		return nil, err
	}

	data := []byte(*jsn)
	addr := fmt.Sprintf("%s%s.json", ZooplaListingUrl, method)

	req, err := http.NewRequest("POST", addr, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return nil, err
	}

	headerValue := fmt.Sprintf("application/json; profile=%s%s.json", ZooplaListingHeaderUrl, method)
	req.Header.Set("Content-Type", headerValue)

	respBytes, err := getBytes(req, ZooplaClientTimeout)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}
