package datastore

import (
	"context"
	"time"

	"content_collector/internal/infrastructure/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClient struct {
	Client *mongo.Client
	logger logger.Logger
}

func NewClientMongoDB(uri string, user string, password string, localLogger logger.Logger) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	credentials := options.Credential{
		Username: user,
		Password: password,
	}
	clientOptions := options.Client().ApplyURI(uri).SetAuth(credentials)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return &MongoClient{
		Client: client,
		logger: localLogger,
	}, nil
}

func (m *MongoClient) Disconnect() {
	err := m.Client.Disconnect(context.Background())
	if err != nil {
		m.logger.Fatal(err)
	}
}

func (m *MongoClient) GetClient() *mongo.Client {
	return m.Client
}
