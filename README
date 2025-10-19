# YouTube Downloader API

A REST API built with Go for downloading YouTube videos, using Echo framework and kkdai/youtube library.

## 🚀 Features

- **Get video information**: Lists all available formats (quality, codec, bitrate)
- **Download videos**: Streams video directly to the client

## 🛠️ Tech Stack

- **Go 1.24.4+**
- **Echo v4** - Web framework
- **kkdai/youtube/v2** - YouTube interaction library

## 📁 Project Structure

```
my-downloader/
├── cmd/
│   └── main.go              # Application entry point
├── handler/
│   └── video_handler.go     # HTTP handlers
├── services/
│   └── youtube.go           # YouTube business logic
├── models/
│   └── video.go             # Data models
├── go.mod
├── go.sum
└── Dockerfile
```

## 🚀 Running Locally

### Prerequisites

- Go 1.24.4 or higher
- Git

### Installation

1. Clone the repository:

```bash
git clone https://github.com/your-username/youtube-downloader-api.git
cd youtube-downloader-api
```

2. Install dependencies:

```bash
go mod download
```

3. Run the server:

```bash
go run cmd/main.go
```

The server will be running at `http://localhost:8080`

## 📡 API Endpoints

### 1. Get Video Information

**POST** `/api/video/info`

Returns detailed information about the video and available formats.

**Request Body:**

```json
{
  "url": "https://www.youtube.com/watch?v=VIDEO_ID"
}
```

**Response (200 OK):**

```json
{
  "title": "Video Title",
  "author": "YouTube Channel",
  "duration": 180,
  "formats": [
    {
      "itag": 18,
      "quality": "360p",
      "mime_type": "video/mp4; codecs=\"avc1.42001E, mp4a.40.2\"",
      "bitrate": 500000
    },
    {
      "itag": 22,
      "quality": "720p",
      "mime_type": "video/mp4; codecs=\"avc1.64001F, mp4a.40.2\"",
      "bitrate": 2000000
    }
  ]
}
```

### 2. Download Video

**GET** `/api/video/download?url={VIDEO_URL}&itag={ITAG}`

Downloads/streams the video in the specified format.

**Query Parameters:**

- `url` (string, required): Full YouTube video URL
- `itag` (int, required): Format code obtained from `/info` endpoint

**Example:**

```
GET /api/video/download?url=https://www.youtube.com/watch?v=dQw4w9WgXcQ&itag=18
```

**Response:**

- Video file streaming
- Headers: Appropriate `Content-Type` and `Content-Disposition`

## 🧪 Testing the API

### Using cURL

**1. Get video info:**

```bash
curl -X POST http://localhost:8080/api/video/info \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ"}'
```

**2. Download video:**

```bash
curl "http://localhost:8080/api/video/download?url=https://www.youtube.com/watch?v=dQw4w9WgXcQ&itag=18" \
  --output video.mp4
```

### Using Browser

To test the download, simply open in your browser:

```
http://localhost:8080/api/video/download?url=https://www.youtube.com/watch?v=dQw4w9WgXcQ&itag=18
```

## 🐳 Docker

### Build image:

```bash
docker build -t youtube-downloader-api .
```

### Run container:

```bash
docker run -p 8080:8080 youtube-downloader-api
```

## ☁️ Deployment

### Render

1. Create an account at [render.com](https://render.com)
2. Connect your GitHub repository
3. Configure:
   - **Environment:** Go
   - **Build Command:** `go build -o bin/server ./cmd/main.go`
   - **Start Command:** `./bin/server`
4. Deploy!

**Important:** Make sure your `main.go` reads the `PORT` environment variable:

```go
port := os.Getenv("PORT")
if port == "" {
    port = "8080"
}
e.Start(":" + port)
```

### Fly.io

```bash
# Install CLI
curl -L https://fly.io/install.sh | sh

# Login
flyctl auth login

# Deploy
flyctl launch
flyctl deploy
```

## ⚙️ Environment Variables

- `PORT`: Port where the server will run (default: 8080)

## 🔒 CORS

The API comes pre-configured with CORS enabled to accept requests from any origin.

## 📝 Important Notes

- The API streams directly from YouTube, no videos are stored on the server
- First request after inactivity may take ~30s (if hosted on Render free tier)
- Invalid itag formats will return 400 error
- The server automatically hibernates after 15 minutes of inactivity on free hosting tiers

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License.

## ⚠️ Legal Disclaimer

This API is for educational purposes only. Please respect YouTube's Terms of Service and content creators' copyrights.

---

Made with ❤️ using Go
