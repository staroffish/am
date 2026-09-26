package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/service"
	"github.com/staroffish/am/internal/util"
)

type TaskWithAnime struct {
	model.DownloadTask
	Name      string `json:"name"`
	StorePath string `json:"store_path"`
}

type TaskHandler struct {
	svc      *service.DownloadManagerService
	animeSvc *service.AnimeService
	log      *log.Logger
}

func NewTaskHandler(svc *service.DownloadManagerService, animeSvc *service.AnimeService, logger *log.Logger) *TaskHandler {
	return &TaskHandler{svc: svc, animeSvc: animeSvc, log: logger}
}

func (h *TaskHandler) Register(e *echo.Group) {
	e.GET("", h.List)
	e.POST("", h.Create)
	e.GET("/:id", h.Get)
	e.PUT("/:id", h.Update)
	e.DELETE("/:id", h.Delete)
	e.POST("/scan", h.Scan)
	e.POST("/scan-and-download", h.ScanAndDownload)
}

func (h *TaskHandler) List(c echo.Context) error {
	tasks, err := h.svc.ListTasks(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	result := make([]TaskWithAnime, 0, len(tasks))
	for _, t := range tasks {
		r := TaskWithAnime{DownloadTask: t}
		if a, err := h.animeSvc.Get(c.Request().Context(), t.AnimeID); err == nil {
			r.Name = a.AnimeNameJp
			r.StorePath = a.StorDir
		}
		result = append(result, r)
	}
	return c.JSON(http.StatusOK, result)
}

func (h *TaskHandler) Get(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	task, err := h.svc.GetTask(c.Request().Context(), int32(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "task not found")
	}
	r := TaskWithAnime{DownloadTask: *task}
	if a, err := h.animeSvc.Get(c.Request().Context(), task.AnimeID); err == nil {
		r.Name = a.AnimeNameJp
		r.StorePath = a.StorDir
	}
	return c.JSON(http.StatusOK, r)
}

func (h *TaskHandler) Create(c echo.Context) error {
	var req struct {
		Regexp        string `json:"regexp"`
		LatestChapter int32  `json:"latest_chapter"`
		AnimeID       string `json:"anime_id"`
		Name          string `json:"name"`
		StorePath     string `json:"store_path"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	task := model.DownloadTask{
		Regexp:        req.Regexp,
		LatestChapter: req.LatestChapter,
		AnimeID:       req.AnimeID,
	}
	if err := h.svc.AddTask(c.Request().Context(), &task); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	_, err := h.animeSvc.Get(c.Request().Context(), task.AnimeID)
	if err != nil {
		now := time.Now()
		y, s := util.NearestSeason(now)
		ani := &model.Anime{
			AnimeNameJp: req.Name,
			Status:      "连载中",
			SerialsDuri: fmt.Sprintf("%04d/%02d~", y, s),
			StorDir:     req.StorePath,
			CreatedAt:   now,
			UpdatedAt:   now,
		}
		ani.ID = task.AnimeID
		h.animeSvc.Save(c.Request().Context(), ani)
	}

	r := TaskWithAnime{DownloadTask: task, Name: req.Name, StorePath: req.StorePath}
	return c.JSON(http.StatusOK, r)
}

func (h *TaskHandler) Update(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	var task model.DownloadTask
	if err := c.Bind(&task); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	task.ID = int32(id)
	if err := h.svc.UpdateTask(c.Request().Context(), &task); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Delete(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 32)

	task, err := h.svc.GetTask(c.Request().Context(), int32(id))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "task not found")
	}

	if err := h.svc.DeleteTask(c.Request().Context(), int32(id)); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if task.AnimeID != "" {
		if err := h.animeSvc.MarkDone(c.Request().Context(), task.AnimeID); err != nil {
			h.log.Printf("mark anime %s done after task %d delete: %v", task.AnimeID, id, err)
		}
	}

	return c.NoContent(http.StatusOK)
}

func (h *TaskHandler) Scan(c echo.Context) error {
	tasks, err := h.svc.Scan(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if tasks == nil {
		tasks = []model.MatchedTask{}
	}
	return c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) ScanAndDownload(c echo.Context) error {
	tasks, err := h.svc.ScanAndDownload(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if tasks == nil {
		tasks = []model.MatchedTask{}
	}
	return c.JSON(http.StatusOK, tasks)
}
