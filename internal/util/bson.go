package util

import (
	"go.mongodb.org/mongo-driver/bson"
)

func ConvertToBSON(search map[string]string) bson.D {
	filter := bson.D{}
	for k, v := range search {
		filter = append(filter, bson.E{Key: k, Value: v})
	}
	return filter
}
