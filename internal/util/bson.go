package util

import (
	"go.mongodb.org/mongo-driver/bson"
)

func ConvertToBSON(search map[string]any) bson.D {
	filter := bson.D{}
	for k, v := range search {
		filter = append(filter, bson.E{Key: k, Value: v})
	}
	return filter
}

func StructToMap(s any) (map[string]any, error) {
	data, err := bson.Marshal(s)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	err = bson.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
