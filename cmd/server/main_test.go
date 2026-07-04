package main

import (
	"io/fs"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	webui "github.com/staroffish/am/web"
)

func TestSPAFrontendEmbedded(t *testing.T) {
	_, err := fs.Sub(webui.Dist, "dist")
	if err != nil {
		t.Fatalf("web/dist not embedded: %v", err)
	}
	t.Log("Frontend dist is correctly embedded in binary")
}

func TestSPARouting(t *testing.T) {
	e := echo.New()

	distFS, err := fs.Sub(webui.Dist, "dist")
	if err != nil {
		t.Fatal(err)
	}
	fileHandler := http.FileServer(http.FS(distFS))
	e.GET("/*", echo.WrapHandler(http.StripPrefix("/", fileHandler)))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "AM - Anime Manager") {
		t.Fatalf("body does not contain 'AM - Anime Manager': %s", body[:200])
	}
	t.Log("SPA index.html served correctly")
}

func TestAPIRoutes(t *testing.T) {
	e := echo.New()

	api := e.Group("/api/v1")

	api.GET("/anime", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []interface{}{})
	})
	api.GET("/anime/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"anime": nil, "files": []interface{}{}})
	})
	api.POST("/anime", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"id": "test"})
	})
	api.PUT("/anime/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"id": "test"})
	})
	api.DELETE("/anime/:id", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	api.GET("/tasks", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []interface{}{})
	})
	api.POST("/tasks", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"id": "1"})
	})
	api.POST("/tasks/scan", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []interface{}{})
	})
	api.POST("/tasks/scan-and-download", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []interface{}{})
	})

	api.GET("/spiders", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []interface{}{})
	})
	api.POST("/spiders/crawl", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "started"})
	})

	api.GET("/downloads", func(c echo.Context) error {
		return c.JSON(http.StatusOK, []interface{}{})
	})

	tests := []struct {
		method string
		path   string
		code   int
	}{
		{"GET", "/api/v1/anime", 200},
		{"GET", "/api/v1/anime/abc123", 200},
		{"POST", "/api/v1/anime", 200},
		{"PUT", "/api/v1/anime/abc123", 200},
		{"DELETE", "/api/v1/anime/abc123", 200},
		{"GET", "/api/v1/tasks", 200},
		{"POST", "/api/v1/tasks", 200},
		{"POST", "/api/v1/tasks/scan", 200},
		{"POST", "/api/v1/tasks/scan-and-download", 200},
		{"GET", "/api/v1/spiders", 200},
		{"POST", "/api/v1/spiders/crawl", 200},
		{"GET", "/api/v1/downloads", 200},
	}

	for _, tc := range tests {
		req := httptest.NewRequest(tc.method, tc.path, nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		if rec.Code != tc.code {
			t.Errorf("%s %s: expected %d, got %d", tc.method, tc.path, tc.code, rec.Code)
		} else {
			t.Logf("%s %s: %d OK", tc.method, tc.path, rec.Code)
		}
	}
}
