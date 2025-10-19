package handler

import (
	"api-yt-downloader/models"
	"api-yt-downloader/services"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type VideoHandler struct {
	youtubeService *services.YouTubeService
}

func NewVideoHandler(ys *services.YouTubeService) *VideoHandler {
	return &VideoHandler{
		youtubeService: ys,
	}
}

func (h *VideoHandler) GetInfo(c echo.Context) error {
	var req models.DownloadRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	url := req.URL
	yts, err := h.youtubeService.GetVideoInfo(url)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, yts)
}

func (h *VideoHandler) Download(c echo.Context) error {
	url := c.QueryParam("url")
	itagStr := c.QueryParam("itag")

	itagConvert, err := strconv.Atoi(itagStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	video, format, err := h.youtubeService.DownloadVideo(url, itagConvert)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	c.Response().Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, video.Title))
	c.Response().Header().Set("Content-Type", format.MimeType)

	readCloser, _, err := h.youtubeService.GetStream(video, format)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	_, err = io.Copy(c.Response(), readCloser)
	if err != nil {
		return err
	}
	readCloser.Close()
	// fmt.Printf("Download file sucess: %s\n", written)
	// log.Printf("Copied %d bytes from source.txt to destination.txt", written)

	return nil

}
