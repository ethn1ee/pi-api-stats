package stats

import (
	"context"
	"errors"

	"github.com/shirou/gopsutil/v3/net"
)

type NetworkStat struct {
	BytesIn    uint64
	BytesOut   uint64
	PacketsIn  uint64
	PacketsOut uint64
}

func GetNetworkStat(ctx context.Context) (NetworkStat, error) {
	io, err := net.IOCountersWithContext(ctx, false)
	if err != nil {
		return NetworkStat{}, err
	}
	if len(io) < 1 {
		return NetworkStat{}, errors.New("empty slice returned")
	}

	return NetworkStat{
		BytesIn:    io[0].BytesRecv,
		BytesOut:   io[0].BytesSent,
		PacketsIn:  io[0].PacketsRecv,
		PacketsOut: io[0].PacketsSent,
	}, nil
}
