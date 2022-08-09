// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"

	//"github.com/TIBCOSoftware/labs-air/edgexfoundry/device-sound-simulator/driver"
	"github.com/TIBCOSoftware/labs-air-edgex/device-sound-simulator/driver"

	"github.com/edgexfoundry/device-sdk-go/v2/pkg/startup"
)

const (
	serviceName string = "edgex-sound-simulator"
)

func main() {
	os.Setenv("EDGEX_SECURITY_SECRET_STORE", "false")
	sd := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, device_sound_simulator.Version, sd)
}
