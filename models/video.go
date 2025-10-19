package models

import "time"

type DownloadRequest struct {
	URL string `json:"url" validate:required`
}

type VideoInfo struct {
	Title    string        `json:"title"`
	Author   string        `json:"author"`
	Duration time.Duration `json:"duration"`
	Formats  []Format      `json:"formats"`
}

type Format struct {
	Itag     int    `json:itag`
	Quality  string `json:quality`
	MimeType string `json:mimeType`
	Bitrate  int    `json:bitrate`
}
