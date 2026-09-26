package handler

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/store"
)

type MagnetHandler struct {
	db  *store.DB
	log *log.Logger
}

func NewMagnetHandler(db *store.DB, logger *log.Logger) *MagnetHandler {
	return &MagnetHandler{db: db, log: logger}
}

func (h *MagnetHandler) Register(e *echo.Group) {
	e.GET("", h.List)
	e.GET("/dates", h.Dates)
}

func (h *MagnetHandler) Dates(c echo.Context) error {
	dates, err := h.db.ListMagnetDates(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, dates)
}

func (h *MagnetHandler) List(c echo.Context) error {
	date := c.QueryParam("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	keyword := strings.ToLower(c.QueryParam("keyword"))

	rows, err := h.db.ListMagnetsByDate(c.Request().Context(), date)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result := make([]model.AnimeMagnet, 0, len(rows))
	for _, m := range rows {
		if keyword != "" && !strings.Contains(strings.ToLower(m.Name), keyword) {
			continue
		}
		result = append(result, model.AnimeMagnet{
			Name:       m.Name,
			MagnetLink: m.MagnetLink,
		})
	}

	return c.JSON(http.StatusOK, result)
}
