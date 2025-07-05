# ---- Generate Stage ----
FROM ghcr.io/a-h/templ:latest AS generate-stage
WORKDIR /app
COPY --chown=65532:65532 . .
RUN ["templ", "generate"]

# ---- Build Stage ----
FROM golang:1.24.2-alpine AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOARCH=arm64 CGO_ENABLED=0 GOOS=linux go build -o /app/myapp ./cmd/web/

# --- DEBUG STEP 1 ---
# Check if the 'myapp' binary exists after the build.
RUN ls -la /app

# ---- Deploy Stage ----
FROM gcr.io/distroless/base-debian12 AS deploy-stage
WORKDIR /app
COPY --from=build-stage /app/myapp .
COPY --from=build-stage /app/internal/db/migrations ./internal/db/migrations
COPY --from=build-stage /app/ui ./ui

EXPOSE 4000
USER nonroot:nonroot
ENTRYPOINT ["/app/myapp"]
