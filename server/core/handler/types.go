package handler

import (
	"server/core/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type Response struct {
	Stat    string `json:"stat"`
	Message string `json:"message"`
}

type Handler struct {
	Configs     *config.Configs
	MongoClient *mongo.Client
}

const (
	statSuccess string = "success"
	statError   string = "error"
)
