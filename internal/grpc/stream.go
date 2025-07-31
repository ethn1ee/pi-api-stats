package grpc

import (
	"log"
	pb "pi/stats/proto"
	"time"
)

func (s *Server) StreamAllStats(in *pb.Empty, stream pb.Stats_StreamAllStatsServer) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctx := stream.Context()
			// This could be more efficient by fetching in parallel
			cpu, err := s.GetCpu(ctx, in)
			if err != nil {
				log.Printf("failed to get cpu stats: %v", err)
				continue
			}
			disk, err := s.GetDisk(ctx, in)
			if err != nil {
				log.Printf("failed to get disk stats: %v", err)
				continue
			}
			host, err := s.GetHost(ctx, in)
			if err != nil {
				log.Printf("failed to get host stats: %v", err)
				continue
			}
			memory, err := s.GetMemory(ctx, in)
			if err != nil {
				log.Printf("failed to get memory stats: %v", err)
				continue
			}
			temp, err := s.GetTemperature(ctx, in)
			if err != nil {
				log.Printf("failed to get temperature stats: %v", err)
				continue
			}

			if err := stream.Send(&pb.AllStats{
				Cpu:         cpu,
				Disk:        disk,
				Host:        host,
				Memory:      memory,
				Temperature: temp,
			}); err != nil {
				log.Printf("failed to send stats: %v", err)
				return err
			}
		case <-stream.Context().Done():
			log.Println("client disconnected")
			return nil
		}
	}
}
