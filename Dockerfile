# --- builder ---
FROM golang:1.24.5-alpine AS builder
RUN apk add --no-cache git protobuf-dev
WORKDIR /app
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go generate ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/server ./cmd/main.go

# --- distroless ---
FROM gcr.io/distroless/static-debian12
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 50051
ENTRYPOINT ["./server"]
