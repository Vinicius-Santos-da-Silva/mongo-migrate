package entity

type CollectionSpecification struct {
	Name string `bson:"name"`
	Type string `bson:"type"`
}
