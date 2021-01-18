package base

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func initMongoDriver(database string, collect string) *mongo.Collection {
	client := Conf.MongoDB
	collection := client.Database(database).Collection(collect)
	return collection
}

func GetMongoCollection(database string, collect string) *mongo.Collection {
	return initMongoDriver(database, collect)
}

func AddMongoOne(database string, collect string, content interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := initMongoDriver(database, collect)
	res, err := collection.InsertOne(ctx, content)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateMongoOne(database string, collect string, where interface{}, content interface{}) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := initMongoDriver(database, collect)
	res, err := collection.UpdateOne(ctx, where, bson.M{"$set": content})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func FindMongoOne(database string, collect string, content interface{}) (bson.Raw, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := initMongoDriver(database, collect)
	res := collection.FindOne(ctx, content)
	result, err := res.DecodeBytes()
	if err != nil {
		return nil, err
	}
	return result, nil
}
