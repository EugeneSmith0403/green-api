package main

import (
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/evgenijkuznecov/green-api/docs"
	"github.com/evgenijkuznecov/green-api/internal/config"
	"github.com/evgenijkuznecov/green-api/internal/greenapi"
	"github.com/evgenijkuznecov/green-api/internal/handler"
	"github.com/evgenijkuznecov/green-api/internal/middleware"
	"github.com/evgenijkuznecov/green-api/internal/service"
	"github.com/evgenijkuznecov/green-api/internal/static"
)

// @title           GREEN-API Integration
// @version         1.0
// @description     Proxy server for GREEN-API WhatsApp methods.
// @BasePath        /
func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	client := greenapi.NewClient(cfg.GreenAPIURL)
	svc := service.New(client)
	h := handler.New(svc)

	subFS, err := fs.Sub(static.FS, "dist")
	if err != nil {
		log.Fatalf("static fs: %v", err)
	}

	indexHTML, err := fs.ReadFile(subFS, "index.html")
	if err != nil {
		log.Fatalf("read index.html: %v", err)
	}

	r := gin.New()
	r.Use(middleware.Logger(), middleware.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.RegisterRoutes(r)

	fileServer := http.FileServer(http.FS(subFS))

	r.NoRoute(func(c *gin.Context) {
		p := strings.TrimPrefix(c.Request.URL.Path, "/")

		if _, err := fs.Stat(subFS, p); err == nil {
			fileServer.ServeHTTP(c.Writer, c.Request)
			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", indexHTML)
	})

	log.Printf("server listening on %s", cfg.Addr())
	if err := r.Run(cfg.Addr()); err != nil {
		log.Fatal(err)
	}
}
