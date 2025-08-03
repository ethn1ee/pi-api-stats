# --- deps ---
# This stage is dedicated to downloading dependencies. It will only be re-run
# when go.mod or go.sum changes, caching the expensive dependency download step.
FROM golang:1.24.5-alpine AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# --- builder ---
# This stage builds the application. It uses the cached dependencies from the
# 'deps' stage and only re-runs if the application source code changes.
FROM golang:1.24.5-alpine AS builder
WORKDIR /app
# Copy the cached dependencies from the 'deps' stage.
COPY --from=deps /go/pkg/mod /go/pkg/mod
# Copy the application source code.
COPY . .
# Build the application. The Go compiler will find the dependencies in the
# module cache populated in the previous step.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/server ./cmd/main.go

# --- distroless ---
# The final, minimal image.
FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 50051
ENTRYPOINT ["./server"]