package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"connectrpc.com/grpcreflect"
	"github.com/lmittmann/tint"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/ethn1ee/pi-api-stats/internal/server"
	"github.com/ethn1ee/pi-protos/gen/go/stats/statsconnect"
)

const PORT = 50051

func main() {
	w := os.Stderr

	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	mux := http.NewServeMux()

	// Add reflection
	reflector := grpcreflect.NewStaticReflector(statsconnect.StatsServiceName)
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	// Add stats service
	s := &server.Server{}
	path, handler := statsconnect.NewStatsServiceHandler(s)
	mux.Handle(path, handler)

	// Add CORS middleware
	c := cors.New(cors.Options{
		AllowedMethods: []string{http.MethodGet, http.MethodPost, "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowedOrigins: []string{"*"},
	})
	handlerWithCors := c.Handler(h2c.NewHandler(mux, &http2.Server{}))

	addr := fmt.Sprintf("0.0.0.0:%d", PORT)
	slog.Info("server starting", "addr", addr)

	if err := http.ListenAndServe(addr, handlerWithCors); err != nil {
		slog.Error("failed to listen and serve", slog.Any("error", err))
	}
}
