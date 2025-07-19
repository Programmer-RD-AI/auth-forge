package authforge

import (
	"context"

	"github.com/Programmer-RD-AI/auth-forge/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoDB(ctx context.Context, config *config.MongoConfig) (*mongo.Client, error){
client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(config.Uri))
		if err != nil {
			return nil, err
		}
		return client, nil
}
