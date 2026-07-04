package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/service"
)

type DownloaderHandler struct {
	svc *service.DownloaderService
}

func NewDownloaderHandler(svc *service.DownloaderService) *DownloaderHandler {
	return &DownloaderHandler{svc: svc}
}

func (h *DownloaderHandler) Register(e *echo.Group) {
	e.GET("", h.List)
	e.POST("", h.Add)
	e.DELETE("/:hash", h.Delete)
	e.POST("/:hash/pause", h.Pause)
	e.POST("/:hash/resume", h.Resume)
}

func (h *DownloaderHandler) List(c echo.Context) error {
	list, err := h.svc.List(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if list == nil {
		list = []model.TorrentInfo{}
	}
	return c.JSON(http.StatusOK, list)
}

func (h *DownloaderHandler) Add(c echo.Context) error {
	var req struct {
		Link      string `json:"link"`
		StorePath string `json:"store_path"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := h.svc.Add(c.Request().Context(), req.Link, req.StorePath); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *DownloaderHandler) Delete(c echo.Context) error {
	hash := c.Param("hash")
	if err := h.svc.Delete(c.Request().Context(), hash); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *DownloaderHandler) Pause(c echo.Context) error {
	hash := c.Param("hash")
	if err := h.svc.Pause(c.Request().Context(), hash); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *DownloaderHandler) Resume(c echo.Context) error {
	hash := c.Param("hash")
	if err := h.svc.Resume(c.Request().Context(), hash); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}
