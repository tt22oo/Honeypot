package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"server/core/logger"

	"github.com/labstack/echo/v4"
)

type ReportData struct {
	Time      string `json:"time"`
	IP        string `json:"ip"`
	Action    string `json:"action"`
	Protocol  string `json:"protocol"`
	SessionID string `json:"session_id"`
	Data      string `json:"data"`
}

const logPATH string = "logs.csv"

func (r *ReportData) Save() {
	f, err := os.OpenFile(logPATH, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		data := fmt.Sprintf("Save Error: %s", err.Error())
		logger.Error(data)
		return
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	err = w.Write([]string{r.Time, r.IP, r.Action, r.Protocol, r.Data, r.SessionID})
	if err != nil {
		data := fmt.Sprintf("Write Error: %s", err.Error())
		logger.Error(data)
	}
}

func Report(c echo.Context) error {
	var r ReportData
	err := json.NewDecoder(c.Request().Body).Decode(&r)
	if err != nil {
		return c.JSON(400, Response{
			Stat:    "error",
			Message: fmt.Sprintf("decode error: %s", err.Error()),
		})
	}

	r.Save()
	return c.JSON(200, Response{
		Stat:    "success",
		Message: "added log",
	})
}
