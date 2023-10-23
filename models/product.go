package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id     primitive.ObjectID `bson:"_id"`
	Name   *string            `json:"name"`
	Price  *uint64            `json:"price"`
	Rating *uint8             `json:"rating"`
	Image  *string            `json:"image"`
}
