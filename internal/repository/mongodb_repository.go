package store

import (
	"context"

	"github.com/Programmer-RD-AI/auth-forge/internal/model"
	"github.com/Programmer-RD-AI/auth-forge/internal/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBRepo struct {
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

func (m *MongoDBRepo) Create(ctx context.Context) (interface{}, error) {
	res, err := m.coll.InsertOne(ctx, m.model)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

func (m *MongoDBRepo) Read(ctx context.Context, searchFilter map[string]any, all *bool) ([]any, error) {
	filter := util.ConvertToBSON(searchFilter)
	var res []any
	if *all {
		cur, err := m.coll.Find(ctx, filter)
		if err != nil {
			return nil, err
		}
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

func (m *MongoDBRepo) Update(ctx context.Context, searchFilter map[string]any, updatedModel any, all *bool) (*mongo.UpdateResult, error) {
	// TODO: make it so that the user of this struct can use this function while having a full and then it will read and then check the diff and then only update the news ones
	filter := util.ConvertToBSON(searchFilter)
	update := util.ConvertToBSON(map[string]interface{}{"$set": updatedModel})
	var (
		updatedResponse *mongo.UpdateResult
		err             error
	)
	if *all {
		updatedResponse, err = m.coll.UpdateOne(ctx, filter, update)
	} else {
		updatedResponse, err = m.coll.UpdateMany(ctx, filter, update)
	}
	if err != nil {
		return nil, err
	}
	return updatedResponse, nil
}

func (m *MongoDBRepo) Delete(ctx context.Context, searchFilter map[string]any, all *bool) (*mongo.DeleteResult, error) {
	filter := util.ConvertToBSON(searchFilter)
	var (
		deleteResponse *mongo.DeleteResult
		err            error
	)
	if *all {
		deleteResponse, err = m.coll.DeleteMany(ctx, filter)
	} else {
		deleteResponse, err = m.coll.DeleteOne(ctx, filter)
	}
	return deleteResponse, err
}

func (m *MongoDBRepo) GetByUserId(ctx context.Context, userId string) (*model.Session, error) {
	all := false
	res, err := m.Read(ctx, map[string]any{"userId": userId}, &all)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, mongo.ErrNoDocuments
	}
	session, ok := res[0].(*model.Session)
	if !ok {
		return nil, mongo.ErrNoDocuments
	}
	return session, nil
}

func (m *MongoDBRepo) GetBySessionId(ctx context.Context, sessionId string) (*model.Session, error) {
	all := false
	res, err := m.Read(ctx, map[string]any{"sessionId": sessionId}, &all)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, mongo.ErrNoDocuments
	}
	session, ok := res[0].(*model.Session)
	if !ok {
		return nil, mongo.ErrNoDocuments
	}
	return session, nil
}

func (m *MongoDBRepo) DeleteSessionId(ctx context.Context, sessionId string) bool {
	all := false
	res, err := m.Delete(ctx, map[string]any{"sessionId": sessionId}, &all)
	if err != nil {
		return false
	}
	if res.DeletedCount > 0 {
		return true
	}
	return false
}
