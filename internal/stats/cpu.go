package stats

import (
	"context"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CpuStat struct {
	Percent []float64
	Clock   float64
}

func GetCpuStat(ctx context.Context) (CpuStat, error) {
	percent, err := cpu.PercentWithContext(ctx, 0, true)
	if err != nil {
		return CpuStat{}, err
	}

	return CpuStat{
		Percent: percent,
		Clock:   cpu.ClocksPerSec,
	}, nil
}
