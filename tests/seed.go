package main

import (
	"fmt"

	repository "github.com/Vinicius-Santos-da-Silva/mongo-migrate/pkg/database_adapter"
	pkg "github.com/Vinicius-Santos-da-Silva/mongo-migrate/pkg/migrate"
	seeds "github.com/Vinicius-Santos-da-Silva/mongo-migrate/tests/seed"
)

func main() {
	fmt.Println("Up seeds...")

	database, err := pkg.MongoConnect("test-migrations")
	pkg.SetDatabase(database)

	if err != nil {
		panic(err)
	}

	repo := repository.NewMigrationRepositoryMongo(database)
	pkg.SetRepository(repo)

	pkg.Register(seeds.NewAddMyIndex(database))
	pkg.Register(seeds.NewAddMyIndex(database))
	pkg.Register(seeds.NewAddMyIndex(database))
	pkg.Register(seeds.NewAddMyIndex(database))
	pkg.Register(seeds.NewAddMyIndex(database))

	if err := pkg.Up(pkg.AllAvailable); err != nil {
		panic(err)
	}

}
