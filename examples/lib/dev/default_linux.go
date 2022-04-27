package dev

import (
	"github.com/TSI-Shoreview/ble"
	"github.com/TSI-Shoreview/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
