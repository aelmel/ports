package mongo

import (
	"context"
	"fmt"
	"github.com/aelmel/ports-infra/port_domain_service/internal/domain"
	"github.com/aelmel/ports-infra/port_domain_service/internal/repository"
	log "github.com/sirupsen/logrus"
	mongodb "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongo struct {
	client     *mongodb.Client
	collection *mongodb.Collection
}

func NewRepo(configuration Configuration) (repository.Repository, error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", configuration.Host, configuration.Port)).SetAuth(options.Credential{
		AuthSource: configuration.Database,
		Username:   configuration.Username,
		Password:   configuration.Password,
	})

	client, err := mongodb.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Error(fmt.Sprintf("Error connectiong to mongodb %v", err))
		return nil, err
	}
	collection := client.Database(configuration.Database).Collection(configuration.Collection)
	return mongo{
		client:     client,
		collection: collection,
	}, nil
}

func (m mongo) InsertPort(ctx context.Context, port domain.Port) error {
	panic("implement me")
}

func (m mongo) GetPort(ctx context.Context, portKey string) (domain.Port, error) {
	panic("implement me")
}

func (m mongo) Close() {
	m.client.Disconnect(context.Background())
}
