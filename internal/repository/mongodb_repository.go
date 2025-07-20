package store

import (
	"context"

	"github.com/Programmer-RD-AI/auth-forge/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBStore struct {
	coll  *mongo.Collection
	model any
}

func HealthCheck(ctx context.Context, client *mongo.Client) bool {
	if err := client.Ping(ctx, nil); err != nil {
		return false
	}
	return true
}

func GetCollection(databaseName string, collectionName string, client *mongo.Client) *mongo.Collection {
	return client.Database(databaseName).Collection(collectionName)
}

func (m *MongoDBStore) Create(ctx context.Context) (interface{}, error) {
	res, err := m.coll.InsertOne(ctx, m.model)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func (m *MongoDBStore) Read(ctx context.Context, searchQuery map[string]string, readAll *bool) ([]any, error) {
	filter := util.ConvertToBSON(searchQuery)
	var res []any
	if *readAll == true {
		cur, err := m.coll.Find(ctx, filter)
		if err = cur.All(ctx, &res); err != nil {
			return nil, err
		}
		defer cur.Close(ctx)
		return res, nil
	}
	var raw_res any
	if err := m.coll.FindOne(ctx, filter).Decode(&raw_res); err != nil {
		return nil, err
	}
	res = append(res, raw_res)
	return res, nil
}
