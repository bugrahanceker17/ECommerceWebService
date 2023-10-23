package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	FirstName      *string            `json:"firstName"  validate:"required,min=2,max=30"`
	LastName       *string            `json:"lastName" validate:"required,min=2,max=30"`
	Password       *string            `json:"password" validate:"required,min=6"`
	Email          *string            `json:"email" validate:"email,required"`
	PhoneNumber    *string            `json:"phoneNumber"`
	Token          *string            `json:"token"`
	RefreshToken   *string            `json:"refreshToken"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
	UserId         string             `json:"userId"`
	UserCart       []ProductUser      `json:"userCart" bson:"userCart"`
	AddressDetails []Address          `json:"address" bson:"address"`
	OrderStatus    []Order            `json:"orders" bson:"orders"`
}
