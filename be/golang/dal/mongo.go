package dal

import (
	"context"
	"k8s-go-gin/config"
	"k8s-go-gin/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDAL struct {
	collection *mongo.Collection
}

func NewMongoDAL(databaseName, collectionName string) *MongoDAL {
	collection := config.GetMongoCollection(databaseName, collectionName)
	return &MongoDAL{collection: collection}
}

func (dal *MongoDAL) GetLinkList() ([]models.LinkData, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := dal.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var links []models.LinkData
	if err = cursor.All(ctx, &links); err != nil {
		return nil, err
	}
	return links, nil
}
