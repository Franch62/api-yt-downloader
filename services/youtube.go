package services

import (
	"api-yt-downloader/models"
	"fmt"
	"io"

	"github.com/kkdai/youtube/v2"
)

type YouTubeService struct {
	client *youtube.Client
}

func NewYouTubeService() *YouTubeService {
	return &YouTubeService{
		client: &youtube.Client{},
	}
}

func (s *YouTubeService) GetVideoInfo(url string) (*models.VideoInfo, error) {
	// Aqui você vai:
	// 1. Usar s.client.GetVideo(url) para pegar informações
	// 2. Converter os formatos para seu modelo
	// 3. Retornar as informações
	video, err := s.client.GetVideo(url)
	if err != nil {
		return nil, err
	}

	// TODO: implementar
	var responseFormats []models.Format
	for _, f := range video.Formats {
		responseFormats = append(responseFormats, models.Format{
			Itag:     f.ItagNo,
			Quality:  f.Quality,
			MimeType: f.MimeType,
			Bitrate:  f.Bitrate,
		})
	}

	return &models.VideoInfo{
		Title:    video.Title,
		Author:   video.Author,
		Duration: video.Duration,
		Formats:  responseFormats,
	}, err
}

func (s *YouTubeService) DownloadVideo(url string, itagFormat int) (*youtube.Video, *youtube.Format, error) {
	// Aqui você vai:
	// 1. Pegar o vídeo
	// 2. Encontrar o formato pelo itag
	// 3. Retornar o vídeo e formato para streaming

	youtubeVideo, err := s.client.GetVideo(url)
	if err != nil {
		return nil, nil, err
	}
	// TODO: implementar
	for _, f := range youtubeVideo.Formats {
		if f.ItagNo == itagFormat {
			return youtubeVideo, &f, nil
		}
	}

	return nil, nil, fmt.Errorf("format not found")
}

func (s *YouTubeService) GetStream(video *youtube.Video, format *youtube.Format) (io.ReadCloser, int64, error) {
	return s.client.GetStream(video, format)
}
