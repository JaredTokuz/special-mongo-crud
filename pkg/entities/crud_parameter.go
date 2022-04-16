package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Book Constructs your Book model under entities.
type CrudParameter struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Filter		interface{} 			 `json:"filter" bson:"filter"`
}
