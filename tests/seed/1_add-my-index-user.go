package seed

import "github.com/Vinicius-Santos-da-Silva/mongo-migrate/src/repository"

type addMyIndexUser struct {
	typing     string
	name       string
	version    uint64
	repository repository.OnlineReviewRepository
}

func NewAddMyIndexUser(repository repository.OnlineReviewRepository) *addMyIndexUser {
	return &addMyIndexUser{
		version:    1,
		name:       "addMyIndexUser",
		typing:     "seed",
		repository: repository,
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
	ami.repository.Insert(&repository.OnlineReview{
		Name: "Golang",
	})
	return nil
}

func (ami *addMyIndexUser) Down() error {
	return nil
}
