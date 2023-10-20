package seed

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type addMyIndexUser struct {
	typing  string
	name    string
	version uint64
	db      *mongo.Database
}

func NewAddMyIndexUser(db *mongo.Database) *addMyIndexUser {
	return &addMyIndexUser{
		version: 1,
		name:    "addMyIndexUser",
		typing:  "seed",
		db:      db,
		//repositoryXXXX: XXXXX
	}
}

func (ami *addMyIndexUser) GetName() string {
	return ami.name
}

func (ami *addMyIndexUser) GetType() string {
	return ami.typing
}

func (ami *addMyIndexUser) GetVersion() uint64 {
	return ami.version
}

func (ami *addMyIndexUser) Up() error {
	opt := options.Index().SetName("my-index2")
	keys := bson.D{{"my-key2", 1}}
	model := mongo.IndexModel{Keys: keys, Options: opt}
	_, err := ami.db.Collection("my-coll2").Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Up:", ami.typing, ": ", ami.name, "executed with success")

	return nil
}

func (ami *addMyIndexUser) Down() error {
	_, err := ami.db.Collection("my-coll2").Indexes().DropOne(context.TODO(), "my-index2")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Down: ", ami.typing, ": ", ami.name, "executed with success")
	return nil
}
