package web

import (
	"fmt"
	"server/core/handler"
	"server/core/logger"

	"github.com/labstack/echo/v4"
)

func Start(listen string) {
	e := echo.New()

	e.POST("/report", handler.Report)

	err := e.Start(listen)
	if err != nil {
		data := fmt.Sprintf("Web Server Error: %s", err.Error())
		logger.Error(data)
	}
}
