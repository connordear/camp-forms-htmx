# --- Development Stage ---
FROM golang:1.24.2 AS development

# Set the working directory inside the container
WORKDIR /app

# Install air for live reloading
# (Ensure you have network access during build if behind proxy)
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
# Note: This will be overlayed by the volume mount during dev,
# but it's good practice for potential image rebuilding.
COPY . .

# Copy the air config file
COPY .air.toml .

# RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/main-prod ./main.go

# Expose the port the app runs on
EXPOSE $PORT

ENV DB_PATH = $DB_PATH

# Command to run air, which will build and run the app
# Air will watch for changes in the mounted volume
CMD ["air", "-c", ".air.toml", "-build.args_bin", "-port=:$PORT"]

