package entity

import "time"

type VersionRecord struct {
	Type        string    `bson:"type"`
	Version     uint64    `bson:"version"`
	Description string    `bson:"description,omitempty"`
	Timestamp   time.Time `bson:"timestamp"`
}
