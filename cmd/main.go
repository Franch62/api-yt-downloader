package main

import (
	"api-yt-downloader/handler"
	"api-yt-downloader/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	youtubeService := services.NewYouTubeService()

	videoHandler := handler.NewVideoHandler(youtubeService)

	apiv1 := e.Group("/api")

	apiv1.POST("/video/info", videoHandler.GetInfo)
	apiv1.GET("/video/download", videoHandler.Download)

	apiv1.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))

	for {

	}
}
