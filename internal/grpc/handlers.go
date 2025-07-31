package grpc

import (
	"context"
	"pi/stats/internal/stats"
	pb "pi/stats/proto"
)

func (s *Server) GetCpu(ctx context.Context, in *pb.Empty) (*pb.CpuResponse, error) {
	data, err := stats.GetCpuStat(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CpuResponse{Data: data}, nil
}

func (s *Server) GetDisk(ctx context.Context, in *pb.Empty) (*pb.DiskResponse, error) {
	data, err := stats.GetDiskStat(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.DiskResponse{
		Total:       data.Total,
		Free:        data.Free,
		Used:        data.Used,
		UsedPercent: data.UsedPercent,
	}, nil
}

func (s *Server) GetHost(ctx context.Context, in *pb.Empty) (*pb.HostResponse, error) {
	data, err := stats.GetHostStat(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.HostResponse{
		BootTime:     data.BootTime,
		Uptime:       data.Uptime,
		Processes:    data.Processes,
		Os:           data.Os,
		Platform:     data.Platform,
		Architecture: data.Architecture,
	}, nil
}

func (s *Server) GetMemory(ctx context.Context, in *pb.Empty) (*pb.MemoryResponse, error) {
	data, err := stats.GetMemoryStat(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.MemoryResponse{
		Total:       data.Total,
		Available:   data.Available,
		Used:        data.Used,
		UsedPercent: data.UsedPercent,
	}, nil
}

func (s *Server) GetTemperature(ctx context.Context, in *pb.Empty) (*pb.TemperatureResponse, error) {
	data, err := stats.GetTemperatureStat(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.TemperatureResponse{
		Cpu:  data.Cpu,
		Nvme: data.Nvme,
	}, nil
}
