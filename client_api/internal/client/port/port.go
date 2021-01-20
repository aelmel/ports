package port

import (
	"context"
	"errors"
	"fmt"
	"github.com/aelmel/ports-infra/client_api/internal/domain"
	pb "github.com/aelmel/ports-infra/client_api/internal/proto"
	"google.golang.org/grpc"
	"log"
)

type Client interface {
	AddPort(ctx context.Context, key string, details domain.PortDetails) error
	GetPort(ctx context.Context, key string) (domain.PortDetails, error)
	Close()
}

type grpcClient struct {
	cc     *grpc.ClientConn
	client pb.PortServiceClient
}

func NewGrpcClient(host, port string) (Client, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), opts)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	client := pb.NewPortServiceClient(cc)

	return &grpcClient{cc: cc, client: client}, nil
}

func (g grpcClient) AddPort(ctx context.Context, key string, details domain.PortDetails) error {
	req := pb.InsertPortRequest{Port: &pb.Port{
		Key: key,
		Details: &pb.Details{
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
		},
	}}

	resp, err := g.client.Add(ctx, &req)
	if err != nil {
		return err
	}

	if resp.Status != 0 {
		return errors.New("error adding port")
	}
	return nil
}

func (g grpcClient) GetPort(ctx context.Context, key string) (domain.PortDetails, error) {
	req := &pb.GetPortRequest{Key: key}
	resp, err := g.client.GetPort(ctx, req)
	if err != nil {
		return domain.PortDetails{}, err
	}

	if resp.Status != 0 {
		return domain.PortDetails{}, errors.New("port not found")
	}
	return domain.PortDetails{
		Name:        resp.Details.Name,
		Coordinates: resp.Details.Coordinates,
		City:        resp.Details.City,
		Province:    resp.Details.Province,
		Country:     resp.Details.Country,
		Alias:       resp.Details.Alias,
		Regions:     resp.Details.Regions,
		Timezone:    resp.Details.Timezone,
		Unlocs:      resp.Details.Unlocs,
		Code:        resp.Details.Code,
	}, nil
}

func (g grpcClient) Close() {
	g.cc.Close()
}
