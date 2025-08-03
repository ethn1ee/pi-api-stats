FROM golang:1.24.5-alpine AS deps
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM golang:1.24.5-alpine AS builder
WORKDIR /app
COPY --from=deps /go/pkg/mod /go/pkg/mod
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/api-stats ./cmd/main.go

FROM gcr.io/distroless/static-debian12 AS runner
WORKDIR /app
COPY --from=builder /app/api-stats .
EXPOSE 50051
ENTRYPOINT ["./api-stats"]
