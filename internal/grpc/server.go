package grpc

import (
	"context"
	"log/slog"

	"pi/api/stats/internal/stats"
	pb "github.com/ethn1ee/pi-protos/gen/go/stats"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedStatsServer
}

// a generic helper to create a streaming RPC for a specific statistic.
func streamStat[T any](
	stream grpc.ServerStream,
	fetchFn func(ctx context.Context) (T, error),
	sendFn func(data T) error,
) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data, err := fetchFn(stream.Context())
			if err != nil {
				slog.Error("failed to fetch stat, terminating stream", slog.Any("error", err))
				return err
			}
			if err := sendFn(data); err != nil {
				slog.Error("failed to send stat", slog.Any("error", err))
				return err
			}
		case <-stream.Context().Done():
			slog.Info("client disconnected")
			return nil
		}
	}
}

func (s *Server) StreamCpu(in *pb.Empty, stream pb.Stats_StreamCpuServer) error {
	return streamStat(stream, stats.GetCpuStat, func(data stats.CpuStat) error {
		return stream.Send(&pb.CpuResponse{
			Percent: data.Percent,
			Clock:   data.Clock,
		})
	})
}

func (s *Server) StreamDisk(in *pb.Empty, stream pb.Stats_StreamDiskServer) error {
	return streamStat(stream, stats.GetDiskStat, func(data stats.DiskStat) error {
		return stream.Send(&pb.DiskResponse{
			Total:       data.Total,
			Free:        data.Free,
			Used:        data.Used,
			UsedPercent: data.UsedPercent,
		})
	})
}

func (s *Server) StreamHost(in *pb.Empty, stream pb.Stats_StreamHostServer) error {
	return streamStat(stream, stats.GetHostStat, func(data stats.HostStat) error {
		return stream.Send(&pb.HostResponse{
			BootTime:     data.BootTime,
			Uptime:       data.Uptime,
			Processes:    data.Processes,
			Os:           data.Os,
			Platform:     data.Platform,
			Architecture: data.Architecture,
		})
	})
}

func (s *Server) StreamMemory(in *pb.Empty, stream pb.Stats_StreamMemoryServer) error {
	return streamStat(stream, stats.GetMemoryStat, func(data stats.MemoryStat) error {
		return stream.Send(&pb.MemoryResponse{
			Total:       data.Total,
			Available:   data.Available,
			Used:        data.Used,
			UsedPercent: data.UsedPercent,
		})
	})
}

func (s *Server) StreamTemperature(in *pb.Empty, stream pb.Stats_StreamTemperatureServer) error {
	return streamStat(stream, stats.GetTemperatureStat, func(data stats.TemperatureStat) error {
		return stream.Send(&pb.TemperatureResponse{
			Cpu:  data.Cpu,
			Nvme: data.Nvme,
		})
	})
}
