package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id      bson.ObjectId `bson:"id"`
	Name    string        `bson:"name"`
	Email   string        `bson:"email"`
	Created time.Time     `bson:"created_at"`
	Updated time.Time     `bson:"updated_at"`
}
