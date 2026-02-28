package database

import (
	"context"
	"server/core/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchReports(c *mongo.Client, cfgs *config.Configs) ([]*Report, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	coll := c.Database(cfgs.DBName).Collection("logs")
	cursor, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var reports []*Report
	for cursor.Next(ctx) {
		var report Report
		err := cursor.Decode(&report)
		if err != nil {
			return nil, err
		}
		reports = append(reports, &report)
	}

	return reports, nil
}
