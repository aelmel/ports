package port

import (
	"context"
	"errors"
	"github.com/aelmel/ports-infra/port_domain_service/internal/domain"
	pb "github.com/aelmel/ports-infra/port_domain_service/internal/proto"
	"github.com/aelmel/ports-infra/port_domain_service/internal/repository"
)

type Service interface {
	Add(ctx context.Context, request *pb.InsertPortRequest) (*pb.InsertPortResponse, error)
	GetPort(ctx context.Context, request *pb.GetPortRequest) (*pb.GetPortResponse, error)
}

type grpcPortService struct {
	repo repository.Repository
}

func NewGrpcPortSvc(repo repository.Repository) pb.PortServiceServer {
	return grpcPortService{
		repo: repo,
	}
}


func (g grpcPortService) Add(ctx context.Context, request *pb.InsertPortRequest) (*pb.InsertPortResponse, error) {
	if request.Port == nil {
		return nil, errors.New("Didn't receive port")
	}
	bsonPort := domain.Port{
		Key:         request.Port.Key,
		Name:        request.Port.Details.Name,
		City:        request.Port.Details.City,
		Country:     request.Port.Details.Country,
		Alias:       request.Port.Details.Alias,
		Regions:     request.Port.Details.Regions,
		Coordinates: request.Port.Details.Coordinates,
		Province:    request.Port.Details.Province,
		Timezone:    request.Port.Details.Timezone,
		Unlocs:      request.Port.Details.Unlocs,
		Code:        request.Port.Details.Code,
	}
	err := g.repo.InsertPort(ctx, bsonPort)
	if err != nil {
		return &pb.InsertPortResponse{Status: 1}, nil
	}
	return &pb.InsertPortResponse{Status: 0}, nil
}

func (g grpcPortService) GetPort(ctx context.Context, request *pb.GetPortRequest) (*pb.GetPortResponse, error) {
	details, err := g.repo.GetPort(ctx, request.Key)
	if err != nil {
		return nil, err
	}

	return &pb.GetPortResponse{Status: 0, Details: &pb.Details{
		Name:        details.Name,
		City:        details.City,
		Country:     details.Country,
		Alias:       details.Alias,
		Regions:     details.Regions,
		Coordinates: details.Coordinates,
		Province:    details.Province,
		Timezone:    details.Timezone,
		Unlocs:      details.Unlocs,
		Code:        details.Code,
	}}, nil
}

