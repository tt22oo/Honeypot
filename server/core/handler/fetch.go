package handler

import (
	"fmt"
	"server/core/database"
	"server/core/logger"

	"github.com/labstack/echo/v4"
)

func (h *Handler) FetchReports(c echo.Context) error {
	if h.Configs.Key != c.Request().Header.Get("key") {
		return c.JSON(401, Response{
			Stat:    statError,
			Message: "invalid key",
		})
	}

	reports, err := database.FetchReports(h.MongoClient, h.Configs)
	if err != nil {
		data := fmt.Sprintf("Database Error: %s", err.Error())
		logger.Error(data)

		return c.JSON(500, Response{
			Stat:    "error",
			Message: "database error",
		})
	}

	return c.JSON(200, Response{
		Stat:    "success",
		Message: reports,
	})
}
