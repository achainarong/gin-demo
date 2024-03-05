package utils

import (
	"fmt"
	"gin-demo/src/config"

	"github.com/grafana/pyroscope-go"
)

func InitProfiler() {
	settings := config.LoadSettings()
	fmt.Println(settings.PyroScopeHost)

	pyroscope.Start(pyroscope.Config{
		ApplicationName: "golang.gin-demo",
		ServerAddress:   settings.PyroScopeHost,
		Logger:          pyroscope.StandardLogger,
		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})
}
