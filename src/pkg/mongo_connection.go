package migrate

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// export MONGO_URL=mongodb://localhost:27017
func MongoConnect(databaseHost string, databaseName string) (*mongo.Database, error) {
	var db *mongo.Database

	uri := fmt.Sprintf(databaseHost + "/" + databaseName)
	opt := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(opt)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	db = client.Database(databaseName)

	return db, nil
}
