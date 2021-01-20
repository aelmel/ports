package mongo

import (
	"context"
	"fmt"
	"github.com/aelmel/ports-infra/port_domain_service/internal/domain"
	"github.com/aelmel/ports-infra/port_domain_service/internal/repository"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
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

func (m mongo) InsertorUpdatePort(ctx context.Context, port domain.Port) error {
	pByte, err := bson.Marshal(port)
	if err != nil {
		return err
	}
	var update bson.M
	err = bson.Unmarshal(pByte, &update)
	if err != nil {
		return err
	}

	filter := bson.M{"key": port.Key}
	findOptions := options.Update()
	findOptions.SetUpsert(true)
	_, err = m.collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: update}}, findOptions)
	if err != nil {
		return err
	}

	return nil
}

func (m mongo) GetPort(ctx context.Context, portKey string) (domain.Port, error) {
	var port domain.Port
	err := m.collection.FindOne(ctx, bson.M{"key": portKey}).Decode(&port)
	if err != nil {
		return domain.Port{}, err
	}

	return port, nil
}

func (m mongo) Close() {
	m.client.Disconnect(context.Background())
}
