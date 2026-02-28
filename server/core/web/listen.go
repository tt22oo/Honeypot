package web

import (
	"fmt"
	"server/core/config"
	"server/core/handler"
	"server/core/logger"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func Start(cfgs *config.Configs, client *mongo.Client) {
	e := echo.New()

	h := handler.Handler{
		MongoClient: client,
		Configs:     cfgs,
	}

	e.POST("/honeypot/report", h.Report)
	e.GET("/honeypot/report/fetch", h.FetchReports)

	err := e.Start(cfgs.Listen)
	if err != nil {
		data := fmt.Sprintf("Web Server Error: %s", err.Error())
		logger.Error(data)
	}
}
