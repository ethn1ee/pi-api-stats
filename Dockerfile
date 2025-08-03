# --- deps ---
FROM golang:1.24.5-alpine AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod vendor

# --- builder ---
FROM golang:1.24.5-alpine AS builder
WORKDIR /app
COPY --from=deps /app/vendor /app/vendor
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o /app/server ./cmd/main.go

# --- distroless ---
# The final, minimal image.
FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 50051
ENTRYPOINT ["./server"]
