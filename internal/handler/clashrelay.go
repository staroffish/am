package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/service"
	"github.com/staroffish/am/internal/store"
)

// ClashRelayHandler 提供「Clash 中转」模块的管理接口与 OpenClash 订阅接口。
type ClashRelayHandler struct {
	svc *service.ClashRelayService
	log *log.Logger
}

func NewClashRelayHandler(svc *service.ClashRelayService, logger *log.Logger) *ClashRelayHandler {
	return &ClashRelayHandler{svc: svc, log: logger}
}

func (h *ClashRelayHandler) Register(e *echo.Group) {
	e.GET("/settings", h.GetSettings)
	e.PUT("/settings", h.SaveSettings)
	e.POST("/pull", h.Pull)
	e.GET("/versions", h.ListVersions)
	e.GET("/versions/:id", h.GetVersion)
	e.POST("/versions", h.CreateVersion)
	e.POST("/versions/:id/activate", h.Activate)
	e.DELETE("/versions/:id", h.DeleteVersion)
	e.GET("/status", h.Status)
	// OpenClash 订阅地址（固定路径、无鉴权、仅内网）。
	e.GET("/sub", h.Sub)
}

func (h *ClashRelayHandler) GetSettings(c echo.Context) error {
	st, err := h.svc.Settings(c.Request().Context())
	if err != nil {
		return clashRelayHTTPError(err)
	}
	return c.JSON(http.StatusOK, st)
}

// clashRelaySettingsPatch 用指针区分"没提交"和"提交了空值"，
// 让「订阅源」和「转换设置」可以各自保存、互不阻塞。
type clashRelaySettingsPatch struct {
	AirportURL           *string `json:"airport_url"`
	AirportUA            *string `json:"airport_ua"`
	VPSHost              *string `json:"vps_host"`
	VPSPort              *int    `json:"vps_port"`
	VPSCipher            *string `json:"vps_cipher"`
	VPSPassword          *string `json:"vps_password"`
	VPSNodeName          *string `json:"vps_node_name"`
	VPSGroupName         *string `json:"vps_group_name"`
	KeepVersions         *int    `json:"keep_versions"`
	RulesText            *string `json:"rules_text"`
	RuleProvidersEnabled *bool   `json:"rule_providers_enabled"`
	RuleProvidersText    *string `json:"rule_providers_text"`
}

func (h *ClashRelayHandler) SaveSettings(c echo.Context) error {
	ctx := c.Request().Context()

	var p clashRelaySettingsPatch
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	current, err := h.svc.Settings(ctx)
	if err != nil {
		return clashRelayHTTPError(err)
	}
	if p.AirportURL != nil {
		current.AirportURL = *p.AirportURL
	}
	if p.AirportUA != nil {
		current.AirportUA = *p.AirportUA
	}
	if p.VPSHost != nil {
		current.VPSHost = *p.VPSHost
	}
	if p.VPSPort != nil {
		current.VPSPort = *p.VPSPort
	}
	if p.VPSCipher != nil {
		current.VPSCipher = *p.VPSCipher
	}
	if p.VPSPassword != nil {
		current.VPSPassword = *p.VPSPassword
	}
	if p.VPSNodeName != nil {
		current.VPSNodeName = *p.VPSNodeName
	}
	if p.VPSGroupName != nil {
		current.VPSGroupName = *p.VPSGroupName
	}
	if p.KeepVersions != nil {
		current.KeepVersions = *p.KeepVersions
	}
	if p.RulesText != nil {
		current.RulesText = *p.RulesText
	}
	if p.RuleProvidersEnabled != nil {
		current.RuleProvidersEnabled = *p.RuleProvidersEnabled
	}
	if p.RuleProvidersText != nil {
		current.RuleProvidersText = *p.RuleProvidersText
	}

	st, err := h.svc.SaveSettings(ctx, current)
	if err != nil {
		return clashRelayHTTPError(err)
	}
	return c.JSON(http.StatusOK, st)
}

func (h *ClashRelayHandler) Pull(c echo.Context) error {
	res, err := h.svc.Pull(c.Request().Context())
	if err != nil {
		var ve *service.ValidationError
		if errors.As(err, &ve) {
			return echo.NewHTTPError(http.StatusBadRequest, ve.Error())
		}
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}
	return c.JSON(http.StatusOK, res)
}

func (h *ClashRelayHandler) ListVersions(c echo.Context) error {
	list, err := h.svc.Versions(c.Request().Context())
	if err != nil {
		return clashRelayHTTPError(err)
	}
	if list == nil {
		list = []model.ClashRelayVersion{}
	}
	return c.JSON(http.StatusOK, list)
}

func (h *ClashRelayHandler) GetVersion(c echo.Context) error {
	id, err := clashRelayID(c)
	if err != nil {
		return err
	}
	v, err := h.svc.Version(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "版本不存在")
	}
	return c.JSON(http.StatusOK, v)
}

func (h *ClashRelayHandler) CreateVersion(c echo.Context) error {
	var req struct {
		Content     string `json:"content"`
		ParentID    *uint  `json:"parent_id"`
		Source      string `json:"source"`
		AirportHash string `json:"airport_hash"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if strings.TrimSpace(req.Content) == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "内容不能为空")
	}
	v, unchanged, err := h.svc.SaveVersion(c.Request().Context(), req.Content, req.ParentID, req.Source, req.AirportHash)
	if err != nil {
		return clashRelayHTTPError(err)
	}
	return c.JSON(http.StatusOK, map[string]any{"version": v, "unchanged": unchanged})
}

func (h *ClashRelayHandler) DeleteVersion(c echo.Context) error {
	id, err := clashRelayID(c)
	if err != nil {
		return err
	}
	if err := h.svc.DeleteVersion(c.Request().Context(), id); err != nil {
		if errors.Is(err, store.ErrClashRelayVersionCurrent) {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusNotFound, "版本不存在或删除失败")
	}
	return c.JSON(http.StatusOK, map[string]any{"ok": true})
}

func (h *ClashRelayHandler) Activate(c echo.Context) error {
	id, err := clashRelayID(c)
	if err != nil {
		return err
	}
	if err := h.svc.Activate(c.Request().Context(), id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "版本不存在或切换失败")
	}
	return c.JSON(http.StatusOK, map[string]any{"ok": true})
}

func (h *ClashRelayHandler) Status(c echo.Context) error {
	st, err := h.svc.Status(c.Request().Context())
	if err != nil {
		return clashRelayHTTPError(err)
	}
	return c.JSON(http.StatusOK, st)
}

func (h *ClashRelayHandler) Sub(c echo.Context) error {
	content, err := h.svc.CurrentContent(c.Request().Context())
	if err != nil {
		return clashRelayHTTPError(err)
	}
	if content == "" {
		return echo.NewHTTPError(http.StatusNotFound, "还没有可用版本，请先在「Clash 中转」页面拉取并生成")
	}
	return c.Blob(http.StatusOK, "text/yaml; charset=utf-8", []byte(content))
}

func clashRelayID(c echo.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil || id == 0 {
		return 0, echo.NewHTTPError(http.StatusBadRequest, "非法的版本号")
	}
	return uint(id), nil
}

func clashRelayHTTPError(err error) error {
	var ve *service.ValidationError
	if errors.As(err, &ve) {
		return echo.NewHTTPError(http.StatusBadRequest, ve.Error())
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
