# --- builder ---
FROM golang:1.24.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
COPY proto ./proto
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o /app/server ./cmd/main.go

# --- distroless ---
FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 50051
ENTRYPOINT ["./server"]
