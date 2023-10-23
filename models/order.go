package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Order struct {
	OrderId       primitive.ObjectID `bson:"_id"`
	OrderCart     []ProductUser      `json:"orderList" bson:"orderList"`
	OrderedAt     time.Time          `json:"orderedAt" bson:"orderedAt"`
	Price         int                `json:"totalPrice" bson:"totalPrice"`
	Discount      *int               `json:"discount" bson:"discount"`
	PaymentMethod Payment            `json:"payment" bson:"payment"`
}
