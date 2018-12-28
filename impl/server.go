package impl

import (
	"context"
	"gitlab.okta-solutions.com/mashroom/backend/common/errs"
	"gitlab.okta-solutions.com/mashroom/backend/common/health"
	"gitlab.okta-solutions.com/mashroom/backend/common/log"
	"gitlab.okta-solutions.com/mashroom/backend/zoopla"
	"gitlab.okta-solutions.com/mashroom/backend/zoopla/version"
	"google.golang.org/grpc"
	"net"
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

func ToZooplaListUpdateRequest(req *zoopla.Property) (*ZooplaListingUpdateRequest, error) {
	detailedDescription := []*DetailedDescription{}
	for _, v := range req.DetailedDescription {
		detailedDescription = append(detailedDescription, &DetailedDescription{Text: v.Text})
	}

	result := &ZooplaListingUpdateRequest{
		BranchReference:  ZooplaBranchReference,
		DisplayAddress:   req.Address.Value,
		RentalTerm:       req.RentalTerm,
		//Category:         "commercial" | "residential" category???,
		ListingReference: req.Id,
		// Pricing: &Pricing{
		// 	RentFrequency:   req.Pricing.RentFrequency,
		// 	CurrencyCode:    req.Pricing.CurrencyCode,
		// 	Price:           req.Pricing.Price,
		// 	TransactionType: req.Pricing.TransactionType,
		// }???,
		Location: &Location{
			CountryCode:          req.CountryCode.Value,
			PostalCode:           req.Postcode.Value,
			//PropertyNumberOrName: req.Location.PropertyNumberOrName,
			StreetName:           req.Address.Value,
			TownOrCity:           req.City.Value,
		},
		PropertyType:      property_type???,
		AvailableFromDate: req.Availability Best practice (rent)???,
		DetailedDescription: detailed_description???,
		Bathrooms:           bathrooms??? Best practice (residential)???,
		FurnishedState:      Best practice (residential rent)???,
		LifeCycleStatus:     life_cycle_status???,
		TotalBedrooms:        Best practice (residential)???,
		LivingRooms: Best practice (residential)???,
	}
	return result, nil
}

func GetListUpdateResponse(resp *ZooplaListingUpdateResponse) *zoopla.ListUpdateResponse {

	result := &zoopla.ListUpdateResponse{
		Status:           resp.Status,
		ListingReference: resp.ListingReference,
		Etag:             resp.ListingEtag,
		Url:              resp.URL,
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

func (server *serverImpl) UpdateProperty(ctx context.Context, request *zoopla.Property) (*zoopla.ListUpdateResponse, error) {
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

func (server *serverImpl) DeleteProperty(ctx context.Context, request *zoopla.ListDeleteRequest) (*zoopla.ListDeleteResponse, error) {
	if request == nil {
		return nil, errs.NilRequest
	}

	req, err := GetZooplaListDeleteRequest(request)

	resp, err := ListingDeleteImpl(*req)
	if err != nil {
		return nil, err
	}

	result := GetListDeleteResponse(resp)
	return result, nil
}

func (server *serverImpl) ListProperty(ctx context.Context, request *zoopla.PropertyRequest) (*zoopla.PropertyResponse, error) {
	if request == nil {
		return nil, errs.NilRequest
	}
}

func (server *serverImpl) GetProperty(context.Context, *zoopla.PropertyRequest) (*zoopla.PropertyResponse, error) {
	panic("implement me")
}

func (server *serverImpl) GetPropertyInfo(context.Context, *zoopla.PropertyRequest) (*zoopla.PropertyResponse, error) {
	panic("implement me")
}

func (server *serverImpl) Example(ctx context.Context, request *zoopla.PropertyRequest) (*zoopla.PropertyResponse, error) {
	if request == nil {
		return nil, errs.NilRequest
	}
	return &zoopla.PropertyResponse{
		Status: "Result of " + request.SomeField,
	}, nil
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
