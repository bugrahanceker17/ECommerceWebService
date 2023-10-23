package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Address struct {
	Id      primitive.ObjectID `bson:"_id"`
	House   *string            `json:"houseName" bson:"houseName"`
	Street  *string            `json:"streetName" bson:"streetName"`
	City    *string            `json:"city" bson:"city"`
	PinCode *string            `json:"pinCode" bson:"pinCode"`
}
