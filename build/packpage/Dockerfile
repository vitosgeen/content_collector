# Build stage
FROM golang:1.20-alpine3.17 AS build-stage

WORKDIR /app

COPY ./../../go.mod ./../../go.sum ./
RUN go mod download

COPY ./../../cmd/ ./cmd/
COPY ./../../internal/ ./internal/
COPY ./../../configs/.env ./configs/
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/collector

# Deploy the application binary into a lean image
FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/configs/.env /configs/
COPY --from=build-stage /app/app /

ENTRYPOINT ["/app"]