package impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	webClientTimeout        = 3000
	ZooplaClientTimeout     = 3000
	ZooplaApiKey            = "qh46j22xzsx79byk84vb9tms"
	ZooplaBranchReference   = "MashroomZooplaBranch"
	ZooplaBranchName        = "Mashroom"
	ZooplaBranchEmail       = "Mashroom@Mashroom"
	ZooplaBranchCountryCode = "gb"
	ZooplaBranchPostalCode  = "a11 1A1"
	ZooplaBranchStreeName   = "Mashroom street"
	ZooplaBranchTownOrCity  = "London"
	ZooplaBranchTelephone   = "02079460184"
	ZooplaBranchWebsite     = "http://www.mashroom.com"
	ZooplaListingUrl        = "https://realtime-listings-api.webservices.zpg.co.uk/sandbox/v1/"
	ZooplaListingHeaderUrl  = "https://realtime-listings.webservices.zpg.co.uk/docs/v1.2/schemas/"
)

type ZooplaMethod int

const (
	Branch_update  ZooplaMethod = 0
	Listing_update ZooplaMethod = 1
	Listing_delete ZooplaMethod = 2
	Listing_list   ZooplaMethod = 3
)

func (z ZooplaMethod) String() string {
	names := [...]string{
		"branch/update",
		"listing/update",
		"listing/delete",
		"listing/list"}

	if z < 0 || z > 3 {
		return ""
	}

	return names[z]
}

func BranchUpdateImpl() (*ZooplaBranchUpdateResponse, error) {
	method := Branch_update

	request := ZooplaBranchUpdateRequest{
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

	var result ZooplaBranchUpdateResponse
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func ListingUpdateImpl(request ZooplaListingUpdateRequest) (*ZooplaListingUpdateResponse, error) {
	method := Listing_update
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

	var result ZooplaListingUpdateResponse
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func ListingDeleteImpl(request ZooplaListingDeleteRequest) (*ZooplaListingDeleteResponse, error) {
	method := Listing_delete
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

	var result ZooplaListingDeleteResponse
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func ListingListImpl(request ZooplaListingListRequest) (*ZooplaListingListResponse, error) {
	method := Listing_list
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

	var result ZooplaListingListResponse
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
