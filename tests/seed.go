package main

import (
	// "context"
	// "fmt"
	// "time"

	// migrate "github.com/Vinicius-Santos-da-Silva/mongo-migrate"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	"fmt"

	pkg "github.com/Vinicius-Santos-da-Silva/mongo-migrate/pkg"
	seeds "github.com/Vinicius-Santos-da-Silva/mongo-migrate/tests/seed"
)

func main() {
	fmt.Println("Up seeds...")

	database, err := pkg.MongoConnect("test-migrations")
	pkg.SetDatabase(database)

	if err != nil {
		panic(err)
	}

	pkg.Register(seeds.NewAddMyIndex(database))
	pkg.Register(seeds.NewAddMyIndex(database))
	pkg.Register(seeds.NewAddMyIndex(database))
	pkg.Register(seeds.NewAddMyIndex(database))
	pkg.Register(seeds.NewAddMyIndex(database))

	if err := pkg.Up(pkg.AllAvailable); err != nil {
		panic(err)
	}

}
