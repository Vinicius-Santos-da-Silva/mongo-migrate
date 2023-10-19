package seed

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type addMyIndex struct {
	typing  string
	name    string
	version uint64
	db      *mongo.Database
}

func NewAddMyIndex(db *mongo.Database) *addMyIndex {
	return &addMyIndex{
		version: 1,
		name:    "addMyIndex",
		typing:  "seed",
		db:      db,
		//repositoryXXXX: XXXXX
	}
}

func (ami *addMyIndex) GetName() string {
	return ami.name
}

func (ami *addMyIndex) GetType() string {
	return ami.typing
}

func (ami *addMyIndex) GetVersion() uint64 {
	return ami.version
}

func (ami *addMyIndex) Up() error {
	fmt.Println("Executing seed test...Register")

	opt := options.Index().SetName("my-index")
	keys := bson.D{{"my-key", 1}}
	model := mongo.IndexModel{Keys: keys, Options: opt}
	_, err := ami.db.Collection("my-coll").Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (ami *addMyIndex) Down() error {
	fmt.Println("Executing seed test...Register")

	_, err := ami.db.Collection("my-coll").Indexes().DropOne(context.TODO(), "my-index")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil

}
