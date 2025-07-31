package stats

import (
	"context"
	"log/slog"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CpuStat []float64

func GetCpuStat(ctx context.Context) (CpuStat, error) {
	return getCpuStat(ctx)
}

func getCpuStat(ctx context.Context) (CpuStat, error) {
	info, err := cpu.InfoWithContext(ctx)
	if err != nil {
		return nil, err
	}

	slog.Info("cpu info", "info", info)

	percent, err := cpu.PercentWithContext(ctx, 0, true)
	if err != nil {
		return nil, err
	}

	return CpuStat(percent), nil
}
