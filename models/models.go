package models

import (
	// Import the bson package primitive so that we can generate the ObjectID
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	// ID will be automatically generated when we insert a new Movie
	// bson can also be represented as json so can define rules for it too
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title   string             `json:"title,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
