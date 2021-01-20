package port

import (
	"context"
	"github.com/aelmel/ports-infra/port_domain_service/internal/domain"
	"github.com/aelmel/ports-infra/port_domain_service/internal/mocks"
	pb "github.com/aelmel/ports-infra/port_domain_service/internal/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGrpcPortService_GetPort(t *testing.T) {
	repoMock := mocks.Repository{}

	port := domain.Port{
		Key:         "PORT_ABC",
		Name:        "PORT_NAME",
		City:        "PORT_CITY",
		Country:     "PORT_COUNTRY",
		Alias:       nil,
		Regions:     nil,
		Coordinates: nil,
		Province:    "",
		Timezone:    "",
		Unlocs:      nil,
		Code:        "",
	}

	repoMock.On("GetPort", mock.Anything, port.Key).Return(port, nil)

	svc := grpcPortService{repo: &repoMock}
	ctx := context.Background()
	resp, err := svc.GetPort(ctx, &pb.GetPortRequest{Key: port.Key})

	assert.Nil(t, err)
	assert.Equal(t, resp.Status, int32(0))
	assert.Equal(t, resp.Details.Name, port.Name)
}

func TestGrpcPortService_Add(t *testing.T) {
	repoMock := mocks.Repository{}

	repoMock.On("InsertPort", mock.Anything, mock.AnythingOfType("domain.Port")).Return(nil)
	svc := grpcPortService{repo: &repoMock}
	ctx := context.Background()
	resp, err := svc.Add(ctx, &pb.InsertPortRequest{Port: &pb.Port{
		Key:     "PORT_NY",
		Details: &pb.Details{
			Name:        "NY_PORT",
			City:        "NY",
			Country:     "USA",
			Alias:       nil,
			Regions:     nil,
			Coordinates: nil,
			Province:    "",
			Timezone:    "",
			Unlocs:      nil,
			Code:        "",
		},
	}})

	assert.Nil(t, err)
	assert.Equal(t, resp.Status, int32(0))
}