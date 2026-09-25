package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/staroffish/am/internal/ai"
	"github.com/staroffish/am/internal/config"
	mycron "github.com/staroffish/am/internal/cron"
	"github.com/staroffish/am/internal/downloader"
	"github.com/staroffish/am/internal/handler"
	"github.com/staroffish/am/internal/service"
	"github.com/staroffish/am/internal/store"
	webui "github.com/staroffish/am/web"
)

var configFile = flag.String("config", "configs/config.yaml", "configuration file path")

func main() {
	flag.Parse()

	cfg, err := config.Load(*configFile)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	logWriter := os.Stderr
	if cfg.Log.File != "" {
		f, err := os.OpenFile(cfg.Log.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("open log file %s: %v", cfg.Log.File, err)
		}
		defer f.Close()
		logWriter = f
	}

	logger := log.New(logWriter, "", log.LstdFlags)
	logger.Printf("AM starting...")

	mongo, err := store.NewMongoClient(cfg.MongoDB)
	if err != nil {
		logger.Fatalf("connect mongo: %v", err)
	}
	defer mongo.Close()

	redis, err := store.NewRedisClient(cfg.Redis)
	if err != nil {
		logger.Fatalf("connect redis: %v", err)
	}
	defer redis.Close()

	qbClient := downloader.NewQBittorrent(cfg.QBittorrent, redis, logger)

	dlSvc := service.NewDownloaderService(qbClient, logger)
	animeSvc := service.NewAnimeService(mongo, logger, cfg.Anime.MainPageCount)
	dmSvc := service.NewDownloadManagerService(mongo, redis, dlSvc, animeSvc, logger, cfg.AutoDownload.MagnetTimeout)
	spiderSvc := service.NewSpiderService(cfg.Spiders, redis, dmSvc, logger, cfg.AutoDownload.MagnetTimeout)

	ticker := mycron.Setup(cfg, spiderSvc, cfg.AutoDownload.Enabled, logger)
	if ticker != nil {
		defer ticker.Stop()
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogMethod:   true,
		LogURI:      true,
		LogLatency:  true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Fprintf(logWriter, "%s %s %s %d %s %.3fms\n",
				time.Now().Format("2006-01-02 15:04:05"),
				v.Method, v.URI, v.Status, v.RemoteIP,
				float64(v.Latency)/float64(time.Millisecond),
			)
			return nil
		},
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		logger.Printf("ERROR %s %s: %v", c.Request().Method, c.Request().URL.Path, err)
		c.JSON(code, map[string]string{"error": err.Error()})
	}

	e.Logger.SetOutput(logWriter)

	api := e.Group("/api/v1")

	animeHandler := handler.NewAnimeHandler(animeSvc, cfg.Anime.MainPageCount)
	animeHandler.Register(api.Group("/anime"))

	taskHandler := handler.NewTaskHandler(dmSvc, animeSvc, logger)
	taskHandler.Register(api.Group("/tasks"))

	spiderHandler := handler.NewSpiderHandler(spiderSvc)
	spiderHandler.Register(api.Group("/spiders"))

	dlHandler := handler.NewDownloaderHandler(dlSvc)
	dlHandler.Register(api.Group("/downloads"))

	magnetHandler := handler.NewMagnetHandler(redis, logger)
	magnetHandler.Register(api.Group("/magnets"))

	aiClient := ai.NewClient(cfg.AI, logger)
	if aiClient != nil {
		logger.Printf("AI client initialized: %s (model: %s)", cfg.AI.BaseURL, cfg.AI.Model)
	} else {
		logger.Printf("AI not configured, /api/v1/ai endpoints will return 501")
	}
	aiHandler := handler.NewAIHandler(aiClient, logger)
	aiHandler.Register(api.Group("/ai"))

	api.GET("/config", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"default_store_dir_prefix": cfg.Anime.DefaultStoreDirPrefix,
		})
	})

	distFS, err := fs.Sub(webui.Dist, "dist")
	if err != nil {
		logger.Printf("web/dist not found, SPA not available: %v", err)
	} else {
		fileHandler := http.FileServer(http.FS(distFS))
		e.GET("/*", echo.WrapHandler(http.StripPrefix("/", fileHandler)))
	}

	logger.Printf("server started on %s", cfg.Server.Addr())
	if err := e.Start(cfg.Server.Addr()); err != nil {
		logger.Fatalf("server error: %v", err)
	}
}
