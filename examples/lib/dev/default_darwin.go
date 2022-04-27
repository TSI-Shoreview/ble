package dev

import (
	"github.com/TSI-Shoreview/ble"
	"github.com/TSI-Shoreview/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
