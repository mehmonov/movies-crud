# Build stage
FROM golang:1.23.7-alpine AS builder

WORKDIR /build

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o main cmd/server/main.go

# Final stage
FROM alpine:3.19

WORKDIR /app

RUN adduser -D -g '' appuser

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /build/main .

RUN chown -R appuser:appuser /app

USER appuser

EXPOSE 8080

CMD ["./main"]