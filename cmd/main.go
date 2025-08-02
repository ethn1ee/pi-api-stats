package main

import (
	"log/slog"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/lmittmann/tint"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpchandler "github.com/ethn1ee/pi-api-stats/internal/grpc"
	pb "github.com/ethn1ee/pi-protos/gen/go/api-stats"
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

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		slog.Error("failed to listen", slog.Any("error", err))
	}

	s := grpc.NewServer()
	pb.RegisterStatsServer(s, &grpchandler.Server{})
	reflection.Register(s)

	slog.Info("Server starting", "port", PORT)
	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve", slog.Any("error", err))
	}
}
