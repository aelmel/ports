package repository

import (
	"context"
	"github.com/aelmel/ports-infra/port_domain_service/internal/domain"
)

type Repository interface {
	InsertorUpdatePort(ctx context.Context, port domain.Port) error
	GetPort(ctx context.Context, portKey string) (domain.Port, error)
	Close()
}
