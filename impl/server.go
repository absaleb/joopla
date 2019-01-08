package impl

import (
	"context"
	"github.com/salebgithub/joopla"
	"github.com/salebgithub/joopla/version"
	"gitlab.okta-solutions.com/mashroom/backend/common/errs"
	"gitlab.okta-solutions.com/mashroom/backend/common/health"
	"gitlab.okta-solutions.com/mashroom/backend/common/log"
	"google.golang.org/grpc"
	"net"

	_ "gitlab.okta-solutions.com/mashroom/backend/zoopla"
	_ "gitlab.okta-solutions.com/mashroom/backend/zoopla/version"
)

type Server interface {
	zoopla.ZooplaServiceServer
	Serve(addr string)
	Background()
}

type serverImpl struct {
}


func GetBranchUpdateResponse(resp *ZooplaBranchUpdateResponse) *zoopla.BranchUpdateResponse {

	result := &zoopla.BranchUpdateResponse{
		Status: resp.Status,
	}

	return result
}

func (server *serverImpl) BranchUpdate(ctx context.Context, request *zoopla.BranchUpdateRequest) (*zoopla.BranchUpdateResponse, error) {
	if request == nil {
		return nil, errs.NilRequest
	}

	resp, err := BranchUpdateImpl()
	if err != nil {
		return nil, err
	}

	result := GetBranchUpdateResponse(resp)
	return result, nil
}



func ToZooplaListUpdateRequest(req *zoopla.ListingUpdateRequest) (*ZooplaListingUpdateRequest, error) {
	var detailedDescription []*DetailedDescription
	for _, v := range req.DetailedDescription {
		detailedDescription = append(detailedDescription, &DetailedDescription{Text: v.Text})
	}

	result := &ZooplaListingUpdateRequest{
		BranchReference:  ZooplaBranchReference,
		Category: req.Category,
		ListingReference: req.ListingReference,
		Pricing: &Pricing{
			RentFrequency:   req.Pricing.RentFrequency.String(),
			CurrencyCode:    req.Pricing.CurrencyCode,
			Price:           req.Pricing.Price,
			TransactionType: req.Pricing.TransactionType.String(),
		},
		Location: &Location{
			CountryCode:          req.Location.CountryCode,
			PostalCode:           req.Location.PostalCode,
			PropertyNumberOrName: req.Location.PropertyNumberOrName,
			StreetName:           req.Location.StreetName,
			TownOrCity:           req.Location.TownOrCity,
		},
		PropertyType:   req.PropertyType,
		DetailedDescription: detailedDescription,
		LifeCycleStatus:  req.LifeCycleStatus,
	}
	return result, nil
}

func GetListUpdateResponse(resp *ZooplaListingUpdateResponse) *zoopla.ListingUpdateResponse {

	result := &zoopla.ListingUpdateResponse{
		Status:           resp.Status,
		ListingReference: resp.ListingReference,
		Etag:             resp.ListingEtag,
		Url:              resp.URL,
	}

	return result
}

func (server *serverImpl) UpdateProperty(ctx context.Context, request *zoopla.ListingUpdateRequest) (*zoopla.ListingUpdateResponse, error) {
	if request == nil {
		return nil, errs.NilRequest
	}

	req, err := ToZooplaListUpdateRequest(request)

	resp, err := ListingUpdateImpl(*req)
	if err != nil {
		return nil, err
	}

	result := GetListUpdateResponse(resp)
	return result, nil
}




func ToZooplaListingDeleteRequest(req *zoopla.ListingDeleteRequest) (*ZooplaListingDeleteRequest, error) {
	result := &ZooplaListingDeleteRequest{
		ListingReference: req.ListingReference,
	}
	return result, nil
}

func GetListDeleteResponse(resp *ZooplaListingDeleteResponse) *zoopla.ListingDeleteResponse {

	result := &zoopla.ListingDeleteResponse{
		Status:           resp.Status,
		ListingReference: resp.ListingReference,
	}

	return result
}

func (server *serverImpl) DeleteProperty(ctx context.Context, request *zoopla.ListingDeleteRequest) (*zoopla.ListingDeleteResponse, error) {
	if request == nil {
		return nil, errs.NilRequest
	}

	req, err := ToZooplaListingDeleteRequest(request)

	resp, err := ListingDeleteImpl(*req)
	if err != nil {
		return nil, err
	}

	result := GetListDeleteResponse(resp)
	return result, nil
}



func ToZooplaListingRequest(req *zoopla.ListingListRequest) (*ZooplaListingRequest, error) {
	result := &ZooplaListingRequest{
		BranchReference: req.BranchReference,
	}
	return result, nil
}

func GetListingListResponse(resp *ZooplaListingListResponse) *zoopla.ListingListResponse {

	var listings []*zoopla.Listing
	for _, v := range resp.Listings {
		listings = append(listings, &zoopla.Listing{ListingReference: v.ListingReference, ListingEtag: v.ListingEtag, URL: v.URL})
	}

	result := &zoopla.ListingListResponse{
		Status:           resp.Status,
		BranchReference:resp.BranchReference,
		Listings: listings,
	}

	return result
}

func (server *serverImpl) Listing(ctx context.Context, request *zoopla.ListingListRequest) (*zoopla.ListingListResponse, error) {
	if request == nil {
		return nil, errs.NilRequest
	}

	req, err := ToZooplaListingRequest(request)

	resp, err := ListingListImpl(*req)
	if err != nil {
		return nil, err
	}

	result := GetListingListResponse(resp)
	return result, nil
}



func (server *serverImpl) Background() {
	// background processes
}

func (server *serverImpl) Serve(addr string) {
	if listener, err := net.Listen("tcp", addr); err != nil {
		panic(err)
	} else {
		log.SetHost("zoopla")
		grpcServer := grpc.NewServer()

		zoopla.RegisterZooplaServiceServer(grpcServer, server)

		healthServer := version.NewHealthServer()
		health.RegisterHealthServiceServer(grpcServer, healthServer)

		log.Infoln("zoopla started")
		if err := grpcServer.Serve(listener); err != nil {
			log.Errorln("gRPC error", err)
		}
	}
}

func NewServer() Server {
	server := &serverImpl{}
	go server.Background()
	return server
}
