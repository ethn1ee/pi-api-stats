# Pi Stats API

A gRPC streaming server that provides a real-time monitoring of cpu, disk, memory, etc. every second, utilizing the `gopsutil` package. The server is built for a personal use for monitoring the Raspberry Pi's system resources, but it can be easily adapted for any other platforms.

## Schema

[Protocol Buffer](https://github.com/ethn1ee/pi-protos/blob/main/proto/api-stats/stats.proto)
