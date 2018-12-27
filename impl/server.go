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

func (server *serverImpl) BranchUpdate(ctx context.Context, request *zoopla.BranchUpdateRequest) (*zoopla.BranchUpdateResponse, error) {
	if request == nil {
		return nil, errs.NilRequest
	}

	req, err := GetZooplaBranchUpdateRequest(request)

	resp, err := BranchUpdateImpl(*req)
	if err != nil {
		return nil, err
	}

	result := GetBranchUpdateResponse(resp)
	return result, nil
}

func GetZooplaBranchUpdateRequest(req *zoopla.BranchUpdateRequest) (*ZooplaBranchUpdateRequest, error) {
	result := &ZooplaBranchUpdateRequest{
		BranchName:      req.BranchName,
		BranchReference: req.BranchReference,
		Email:           req.Email,
		Telephone:       req.Telephone,
		Website:         req.Website,
		Location: &Location{
			Coordinates: &Coordinates{
				Longitude: req.Location.Coordinates.Longitude,
				Latitude:  req.Location.Coordinates.Latitude,
			},
			CountryCode:          req.Location.CountryCode,
			County:               req.Location.County,
			Locality:             req.Location.Locality,
			PafUdprn:             req.Location.PafUdprn,
			PostalCode:           req.Location.PostalCode,
			PropertyNumberOrName: req.Location.PropertyNumberOrName,
			StreetName:           req.Location.StreetName,
			TownOrCity:           req.Location.TownOrCity,
			PafAddress: &PafAddress{
				AddressKey:      req.Location.PafAddress.AddressKey,
				OrganisationKey: req.Location.PafAddress.OrganisationKey,
				PostcodeType:    req.Location.PafAddress.PostcodeType,
			},
		},
	}
	return result, nil
}
func GetBranchUpdateResponse(resp *ZooplaBranchUpdateResponse) *zoopla.BranchUpdateResponse {

	result := &zoopla.BranchUpdateResponse{
		Result: resp.result,
	}

	return result
}

func (server *serverImpl) ListProperty(context.Context, *zoopla.PropertyRequest) (*zoopla.PropertyResponse, error) {
	panic("implement me")
}

func (server *serverImpl) UpdateProperty(context.Context, *zoopla.PropertyRequest) (*zoopla.PropertyResponse, error) {
	panic("implement me")
}

func (server *serverImpl) DeleteProperty(context.Context, *zoopla.PropertyRequest) (*zoopla.PropertyResponse, error) {
	panic("implement me")
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
