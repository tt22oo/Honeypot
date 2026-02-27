package database

import (
	"context"
	"server/core/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Report struct {
	Name      string `json:"name" bson:"name"`
	Time      string `json:"time" bson:"time"`
	IP        string `json:"ip" bson:"ip"`
	Action    string `json:"action" bson:"action"`
	Protocol  string `json:"protocol" bson:"protocol"`
	SessionID string `json:"session_id" bson:"session_id"`
	Data      string `json:"data" bson:"data"`
}

func (r *Report) Add(c *mongo.Client, cfgs *config.Configs) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := c.Database(cfgs.DBName).Collection("logs")
	_, err := coll.InsertOne(ctx, r)
	if err != nil {
		return err
	}

	return nil
}
