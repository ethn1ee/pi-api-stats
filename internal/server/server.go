package server

import (
	"context"
	"log/slog"
	"time"

	"connectrpc.com/connect"
	"github.com/ethn1ee/pi-api-stats/internal/stats"
	pb "github.com/ethn1ee/pi-protos/gen/go/api-stats"
	statsconnect "github.com/ethn1ee/pi-protos/gen/go/api-stats/api_statsconnect"
)

type Server struct {
	statsconnect.UnimplementedStatsServiceHandler
}

// a generic helper to create a streaming RPC for a specific statistic.
func streamStat[T any, TRes any](
	ctx context.Context,
	stream *connect.ServerStream[TRes],
	fetchFn func(context.Context) (T, error),
	toResFn func(T) *TRes,
) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			data, err := fetchFn(ctx)
			if err != nil {
				slog.Error("failed to fetch stat, terminating stream", slog.Any("error", err))
				return err
			}
			if err := stream.Send(toResFn(data)); err != nil {
				slog.Error("failed to send stat", slog.Any("error", err))
				return err
			}
		case <-ctx.Done():
			slog.Info("client disconnected")
			return nil
		}
	}
}

func (s *Server) StreamCpu(
	ctx context.Context,
	req *connect.Request[pb.StreamCpuRequest],
	stream *connect.ServerStream[pb.StreamCpuResponse],
) error {
	return streamStat(ctx, stream, stats.GetCpuStat, func(data stats.CpuStat) *pb.StreamCpuResponse {
		return &pb.StreamCpuResponse{
			Percent: data.Percent,
			Clock:   data.Clock,
		}
	})
}

func (s *Server) StreamDisk(
	ctx context.Context,
	req *connect.Request[pb.StreamDiskRequest],
	stream *connect.ServerStream[pb.StreamDiskResponse],
) error {
	return streamStat(ctx, stream, stats.GetDiskStat, func(data stats.DiskStat) *pb.StreamDiskResponse {
		return &pb.StreamDiskResponse{
			Total:       data.Total,
			Free:        data.Free,
			Used:        data.Used,
			UsedPercent: data.UsedPercent,
		}
	})
}

func (s *Server) StreamHost(
	ctx context.Context,
	req *connect.Request[pb.StreamHostRequest],
	stream *connect.ServerStream[pb.StreamHostResponse],
) error {
	return streamStat(ctx, stream, stats.GetHostStat, func(data stats.HostStat) *pb.StreamHostResponse {
		return &pb.StreamHostResponse{
			BootTime:        data.BootTime,
			Uptime:          data.Uptime,
			Processes:       data.Processes,
			Os:              data.Os,
			Platform:        data.Platform,
			PlatformVersion: data.PlatformVersion,
			KernelVersion:   data.KernelVersion,
			Architecture:    data.Architecture,
		}
	})
}

func (s *Server) StreamMemory(
	ctx context.Context,
	req *connect.Request[pb.StreamMemoryRequest],
	stream *connect.ServerStream[pb.StreamMemoryResponse],
) error {
	return streamStat(ctx, stream, stats.GetMemoryStat, func(data stats.MemoryStat) *pb.StreamMemoryResponse {
		return &pb.StreamMemoryResponse{
			Total:       data.Total,
			Available:   data.Available,
			Used:        data.Used,
			UsedPercent: data.UsedPercent,
		}
	})
}

func (s *Server) StreamTemperature(
	ctx context.Context,
	req *connect.Request[pb.StreamTemperatureRequest],
	stream *connect.ServerStream[pb.StreamTemperatureResponse],
) error {
	return streamStat(ctx, stream, stats.GetTemperatureStat, func(data stats.TemperatureStat) *pb.StreamTemperatureResponse {
		return &pb.StreamTemperatureResponse{
			Cpu:  data.Cpu,
			Nvme: data.Nvme,
		}
	})
}

func (s *Server) StreamNetwork(
	ctx context.Context,
	req *connect.Request[pb.StreamNetworkRequest],
	stream *connect.ServerStream[pb.StreamNetworkResponse],
) error {
	return streamStat(ctx, stream, stats.GetNetworkStat, func(data stats.NetworkStat) *pb.StreamNetworkResponse {
		return &pb.StreamNetworkResponse{
			BytesIn:    data.BytesIn,
			BytesOut:   data.BytesOut,
			PacketsIn:  data.PacketsIn,
			PacketsOut: data.PacketsOut,
		}
	})
}
