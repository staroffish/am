package handler

import (
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/service"
)

type AnimeHandler struct {
	svc           *service.AnimeService
	mainPageCount int
}

func NewAnimeHandler(svc *service.AnimeService, mainPageCount int) *AnimeHandler {
	return &AnimeHandler{svc: svc, mainPageCount: mainPageCount}
}

func (h *AnimeHandler) Register(e *echo.Group) {
	e.GET("", h.List)
	e.GET("/:id", h.Get)
	e.GET("/:id/image", h.GetImage)
	e.POST("", h.Create)
	e.PUT("/:id", h.Update)
	e.DELETE("/:id", h.Delete)
	e.PUT("/:id/done", h.MarkDone)
}

func (h *AnimeHandler) List(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	skip := parseIntParam(c, "skip", 0)
	limit := parseIntParam(c, "limit", h.mainPageCount)
	sort := c.QueryParam("sort")
	if sort != "asc" {
		sort = "desc"
	}
	if limit > 200 {
		limit = 200
	}

	list, total, err := h.svc.List(c.Request().Context(), keyword, skip, limit, sort == "asc")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if list == nil {
		list = []model.Anime{}
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"items": list,
		"total": total,
		"limit": limit,
	})
}

func parseIntParam(c echo.Context, name string, defaultVal int) int {
	s := c.QueryParam(name)
	if s == "" {
		return defaultVal
	}
	v, err := strconv.Atoi(s)
	if err != nil || v < 0 {
		return defaultVal
	}
	return v
}

func (h *AnimeHandler) Get(c echo.Context) error {
	id := c.Param("id")
	ani, err := h.svc.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "anime not found")
	}

	files, _ := h.svc.ListFiles(ani.StorDir)
	if files == nil {
		files = []model.AnimeFile{}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"anime":     ani,
		"files":     files,
		"has_image": len(ani.ImageBin) > 0,
	})
}

func (h *AnimeHandler) GetImage(c echo.Context) error {
	id := c.Param("id")
	ani, err := h.svc.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "anime not found")
	}
	if len(ani.ImageBin) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "no image")
	}
	return c.Blob(http.StatusOK, "image/jpeg", ani.ImageBin)
}

func (h *AnimeHandler) Create(c echo.Context) error {
	var req struct {
		AnimeNameCn string `json:"animenamecn"`
		AnimeNameJp string `json:"animenamejp"`
		Cast        string `json:"cast"`
		Type        string `json:"type"`
		Status      string `json:"status"`
		SerialsDuri string `json:"serialsduri"`
		StorDir     string `json:"stordir"`
		ImageURL    string `json:"image_url"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ani := &model.Anime{
		AnimeNameCn: req.AnimeNameCn,
		AnimeNameJp: req.AnimeNameJp,
		Cast:        req.Cast,
		Type:        req.Type,
		Status:      req.Status,
		SerialsDuri: req.SerialsDuri,
		StorDir:     req.StorDir,
	}

	if req.ImageURL != "" {
		resp, err := http.Get(req.ImageURL)
		if err == nil {
			defer resp.Body.Close()
			data, _ := io.ReadAll(resp.Body)
			ani.ImageBin = data
		}
	} 

	if err := h.svc.Save(c.Request().Context(), ani); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, ani)
}

func (h *AnimeHandler) Update(c echo.Context) error {
	id := c.Param("id")

	existing, err := h.svc.Get(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "anime not found")
	}

	var req struct {
		AnimeNameCn string `json:"animenamecn"`
		AnimeNameJp string `json:"animenamejp"`
		Cast        string `json:"cast"`
		Type        string `json:"type"`
		Status      string `json:"status"`
		SerialsDuri string `json:"serialsduri"`
		StorDir     string `json:"stordir"`
		ImageURL    string `json:"image_url"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	existing.AnimeNameCn = req.AnimeNameCn
	existing.AnimeNameJp = req.AnimeNameJp
	existing.Cast = req.Cast
	existing.Type = req.Type
	existing.Status = req.Status
	existing.SerialsDuri = req.SerialsDuri
	existing.StorDir = req.StorDir

	if req.ImageURL != "" {
		resp, err := http.Get(req.ImageURL)
		if err == nil {
			defer resp.Body.Close()
			data, _ := io.ReadAll(resp.Body)
			existing.ImageBin = data
		}
	}

	if err := h.svc.Save(c.Request().Context(), existing); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if id != "" {
		go h.updateRelatedTask(id, existing.StorDir)
	}

	return c.JSON(http.StatusOK, existing)
}

func (h *AnimeHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	if err := h.svc.Delete(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *AnimeHandler) MarkDone(c echo.Context) error {
	id := c.Param("id")
	if err := h.svc.MarkDone(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h *AnimeHandler) updateRelatedTask(animeID, storePath string) {
	_ = animeID
	_ = storePath
}
