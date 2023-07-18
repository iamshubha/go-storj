# Stage 1: Build stage
FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

COPY . .

RUN  CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

# Stage 2: Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/main .
COPY .env .

# Expose port
EXPOSE 8080

CMD ["./main"]
