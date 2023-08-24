FROM golang:1.20

WORKDIR /app
COPY . .
RUN go mod download
RUN go build internal/handler/app/main.go

EXPOSE 8000