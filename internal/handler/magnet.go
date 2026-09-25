package handler

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/store"
)

type MagnetHandler struct {
	redis *store.RedisClient
	log   *log.Logger
}

func NewMagnetHandler(redis *store.RedisClient, logger *log.Logger) *MagnetHandler {
	return &MagnetHandler{redis: redis, log: logger}
}

func (h *MagnetHandler) Register(e *echo.Group) {
	e.GET("", h.List)
	e.GET("/dates", h.Dates)
}

func (h *MagnetHandler) Dates(c echo.Context) error {
	keys, err := h.redis.Keys(c.Request().Context(), "anime:link:*")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	dates := make([]string, 0, len(keys))
	prefix := "anime:link:"
	for _, key := range keys {
		date := strings.TrimPrefix(key, prefix)
		if len(date) == 10 { // YYYY-MM-DD
			dates = append(dates, date)
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(dates)))

	return c.JSON(http.StatusOK, dates)
}

func (h *MagnetHandler) List(c echo.Context) error {
	date := c.QueryParam("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	keyword := c.QueryParam("keyword")

	key := fmt.Sprintf("anime:link:%s", date)
	magnets, err := h.redis.HGetAll(c.Request().Context(), key)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result := make([]model.AnimeMagnet, 0, len(magnets))
	for name, link := range magnets {
		if keyword != "" && !strings.Contains(strings.ToLower(name), strings.ToLower(keyword)) {
			continue
		}
		result = append(result, model.AnimeMagnet{
			Name:       name,
			MagnetLink: link,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return c.JSON(http.StatusOK, result)
}
