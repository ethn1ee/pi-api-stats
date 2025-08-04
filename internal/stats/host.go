package stats

import (
	"context"

	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
)

type HostStat struct {
	BootTime        uint64 `json:"bootTime"`
	Uptime          uint64 `json:"uptime"`
	Processes       uint64 `json:"processes"`
	Os              string `json:"os"`
	Platform        string `json:"platform"`
	PlatformVersion string `json:"platformVersion"`
	KernelVersion   string `json:"kernelVersion"`
	Architecture    string `json:"architecture"`
}

func GetHostStat(ctx context.Context) (HostStat, error) {
	host.EnableBootTimeCache(true)
	infoStat, err := host.InfoWithContext(ctx)
	if err != nil {
		return HostStat{}, err
	}

	load, err := load.MiscWithContext(ctx)
	if err != nil {
		return HostStat{}, err
	}

	return HostStat{
		BootTime:        infoStat.BootTime,
		Uptime:          infoStat.Uptime,
		Processes:       uint64(load.ProcsTotal),
		Os:              infoStat.OS,
		Platform:        infoStat.Platform,
		PlatformVersion: infoStat.PlatformVersion,
		KernelVersion:   infoStat.KernelVersion,
		Architecture:    infoStat.KernelArch,
	}, nil
}
