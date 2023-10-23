package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ProductUser struct {
	ProductId   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"productName" bson:"productName"`
	Price       *uint64            `json:"price" bson:"price"`
	Rating      *uint8             `json:"rating" bson:"rating"`
	Image       *string            `json:"image" bson:"image"`
}
