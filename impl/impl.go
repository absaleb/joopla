package impl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	webClientTimeout       = 3000
	ZooplaClientTimeout    = 3000
	ZooplaApiKey           = "qh46j22xzsx79byk84vb9tms"
	ZooplaListingUrl       = "https://realtime-listings-api.webservices.zpg.co.uk/sandbox/v1/"
	ZooplaListingHeaderUrl = "https://realtime-listings.webservices.zpg.co.uk/docs/v1.2/schemas/"
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

type ZooplaBranchUpdateResponse struct {
	result string
}

func BranchUpdateImpl(request ZooplaBranchUpdateRequest) (*ZooplaBranchUpdateResponse, error) {
	method := Branch_update
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

	client := &http.Client{}
	client.Timeout = ZooplaClientTimeout * time.Millisecond

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &ZooplaBranchUpdateResponse{result: string(body)}, nil
}
