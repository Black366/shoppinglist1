package models1

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id      primitive.ObjectID
	Name    string
	Checked bool
}
