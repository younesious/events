FROM golang:1.22.3-alpine as builder

RUN apk add --no-cache gcc musl-dev
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o main .

# Start a new stage from scratch
FROM gcr.io/distroless/static

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
