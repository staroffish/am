package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/staroffish/am/internal/service"
)

type SpiderHandler struct {
	svc *service.SpiderService
}

func NewSpiderHandler(svc *service.SpiderService) *SpiderHandler {
	return &SpiderHandler{svc: svc}
}

func (h *SpiderHandler) Register(e *echo.Group) {
	e.GET("", h.List)
	e.POST("/crawl", h.CrawlAll)
	e.POST("/:name/crawl", h.CrawlOne)
}

func (h *SpiderHandler) List(c echo.Context) error {
	return c.JSON(http.StatusOK, h.svc.GetStatus())
}

func (h *SpiderHandler) CrawlAll(c echo.Context) error {
	go h.svc.CrawlAll(context.Background())
	return c.JSON(http.StatusOK, map[string]string{"status": "started"})
}

func (h *SpiderHandler) CrawlOne(c echo.Context) error {
	name := c.Param("name")
	go h.svc.Crawl(context.Background(), name)
	return c.JSON(http.StatusOK, map[string]string{"status": "started", "name": name})
}
