package main

import (
	"fmt"

	pkg "github.com/Vinicius-Santos-da-Silva/mongo-migrate/src/pkg"
	repository "github.com/Vinicius-Santos-da-Silva/mongo-migrate/src/repository"
	seeds "github.com/Vinicius-Santos-da-Silva/mongo-migrate/tests/seed"
)

func main() {
	fmt.Println("Up seeds...")

	database, err := pkg.MongoConnect("test-migrations")

	if err != nil {
		panic(err)
	}

	mysqlconn, err := repository.NewConnection()

	if err != nil {
		panic(err)
	}

	defer mysqlconn.Close()

	// migrationRepo := repository.NewMigrationRepositoryMySQL(mysqlconn)
	migrationRepo := repository.NewMigrationRepositoryMongo(database)
	pkg.SetRepository(migrationRepo)

	onlineReviewRepo := repository.NewOnlineRepositoryMongo(database)

	// pkg.Register(seeds.NewAddMyIndex(database))
	pkg.Register(seeds.NewAddMyIndexUser(onlineReviewRepo))

	if err := pkg.Up(pkg.AllAvailable); err != nil {
		panic(err)
	}

}
