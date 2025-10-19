# Stage 1: Build
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copiar go.mod e go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo o código
COPY . .

# Build (apontando para cmd/main.go)
RUN go build -o server ./cmd/main.go

# Stage 2: Run
FROM alpine:latest

WORKDIR /app

# Copiar o binário do stage anterior
COPY --from=builder /app/server .

# Expor porta
EXPOSE 8080

# Rodar o servidor
CMD ["./server"]