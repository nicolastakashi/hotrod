// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.

package main

import (
	"os"

	"github.com/nicolastakashi/hotrod/cmd"
	"github.com/pyroscope-io/client/pyroscope"
)

var customProfiles = []pyroscope.ProfileType{
	pyroscope.ProfileGoroutines,
	pyroscope.ProfileMutexCount,
	pyroscope.ProfileMutexDuration,
	pyroscope.ProfileBlockCount,
	pyroscope.ProfileBlockDuration,
}

func main() {
	pyroscope.Start(pyroscope.Config{
		ApplicationName: os.Getenv("PYROSCOPE_APPLICATION_NAME"),
		ServerAddress:   os.Getenv("PYROSCOPE_SERVER_ADDRESS"),
		Logger:          pyroscope.StandardLogger,
		Tags:            map[string]string{"hostname": os.Getenv("HOSTNAME")},
		ProfileTypes:    append(pyroscope.DefaultProfileTypes, customProfiles...),
	})

	cmd.Execute()
}
