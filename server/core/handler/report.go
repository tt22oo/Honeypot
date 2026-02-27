package handler

import (
	"encoding/json"
	"fmt"
	"server/core/database"
	"server/core/logger"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Report(c echo.Context) error {
	if h.Configs.Key != c.QueryParam("key") {
		return c.JSON(401, Response{
			Stat:    statError,
			Message: "invalid key",
		})
	}

	var r database.Report
	err := json.NewDecoder(c.Request().Body).Decode(&r)
	if err != nil {
		return c.JSON(400, Response{
			Stat:    statError,
			Message: fmt.Sprintf("Decode Error: %s", err.Error()),
		})
	}

	err = r.Add(h.MongoClient, h.Configs)
	if err != nil {
		data := fmt.Sprintf("MongoDB Error: %s", err.Error())
		logger.Error(data)
		return c.JSON(500, Response{
			Stat:    statError,
			Message: "database error",
		})
	}

	return c.JSON(200, Response{
		Stat:    statSuccess,
		Message: "added",
	})
}
