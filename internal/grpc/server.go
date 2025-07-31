package grpc

import (
	pb "pi/stats/proto"
)

type Server struct {
	pb.UnimplementedStatsServer
}
