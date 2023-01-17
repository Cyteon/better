package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
    ID          primitive.ObjectID `json:"id" bson:"_id"`
    Username    string             `json:"username" bson:"username"`
    Content     string             `json:"content" bson:"content"`
    Channel     string             `json:"channel" bson:"channel"`
    Date        string             `json:"date" bson:"date"`
}