package impl

import (
	"encoding/json"
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

	respBytes, err := getBytesByMethod(ZooplaBranchUpdateRequest{}, method)
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

	respBytes, err := getBytesByMethod(request, method)
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

	respBytes, err := getBytesByMethod(request, method)
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

func ListingListImpl(request ZooplaListingRequest) (*ZooplaListingListResponse, error) {
	method := Listing_list

	respBytes, err := getBytesByMethod(request, method)
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
