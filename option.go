package ble

import (
	"time"

	"github.com/melvinto/ble/linux/hci/evt"

	"github.com/melvinto/ble/linux/hci/cmd"
)

// DeviceOption is an interface which the device should implement to allow using configuration options
type DeviceOption interface {
	SetDeviceID(int) error
	SetDialerTimeout(time.Duration) error
	SetListenerTimeout(time.Duration) error
	SetConnParams(cmd.LECreateConnection) error
	SetScanParams(cmd.LESetScanParameters) error
	SetAdvParams(cmd.LESetAdvertisingParameters) error
	SetExitHandler(f func(error)) error
	SetConnectedHandler(f func(evt.LEConnectionComplete)) error
	SetDisconnectedHandler(f func(evt.DisconnectionComplete)) error
	SetPeripheralRole() error
	SetCentralRole() error
	SetTxPoolExtra(int) error
}

// An Option is a configuration function, which configures the device.
type Option func(DeviceOption) error

// OptDeviceID sets HCI device ID.
func OptDeviceID(id int) Option {
	return func(opt DeviceOption) error {
		opt.SetDeviceID(id)
		return nil
	}
}

// OptDialerTimeout sets dialing timeout for Dialer.
func OptDialerTimeout(d time.Duration) Option {
	return func(opt DeviceOption) error {
		opt.SetDialerTimeout(d)
		return nil
	}
}

// OptListenerTimeout sets dialing timeout for Listener.
func OptListenerTimeout(d time.Duration) Option {
	return func(opt DeviceOption) error {
		opt.SetListenerTimeout(d)
		return nil
	}
}

// OptConnParams overrides default connection parameters.
func OptConnParams(param cmd.LECreateConnection) Option {
	return func(opt DeviceOption) error {
		opt.SetConnParams(param)
		return nil
	}
}

// OptScanParams overrides default scanning parameters.
func OptScanParams(param cmd.LESetScanParameters) Option {
	return func(opt DeviceOption) error {
		opt.SetScanParams(param)
		return nil
	}
}

// OptAdvParams overrides default advertising parameters.
func OptAdvParams(param cmd.LESetAdvertisingParameters) Option {
	return func(opt DeviceOption) error {
		opt.SetAdvParams(param)
		return nil
	}
}

func OptExitHandler(f func(error)) Option {
	return func(opt DeviceOption) error {
		opt.SetExitHandler(f)
		return nil
	}
}

func OptConnectHandler(f func(evt.LEConnectionComplete)) Option {
	return func(opt DeviceOption) error {
		opt.SetConnectedHandler(f)
		return nil
	}
}

func OptDisconnectHandler(f func(evt.DisconnectionComplete)) Option {
	return func(opt DeviceOption) error {
		opt.SetDisconnectedHandler(f)
		return nil
	}
}

// OptPeripheralRole configures the device to perform Peripheral tasks.
func OptPeripheralRole() Option {
	return func(opt DeviceOption) error {
		opt.SetPeripheralRole()
		return nil
	}
}

// OptCentralRole configures the device to perform Central tasks.
func OptCentralRole() Option {
	return func(opt DeviceOption) error {
		opt.SetCentralRole()
		return nil
	}
}

// OptTxPoolExtra sets how many extra ACL tx buffers each connection's pool is
// given on top of (bufCnt-1). The default is 10. Set it to 0 on controllers
// with small LE ACL buffer counts (e.g. UART-attached / Nordic) so the host
// respects HCI flow control and does not trigger Data Buffer Overflow
// (HCI event 0x1A).
func OptTxPoolExtra(n int) Option {
	return func(opt DeviceOption) error {
		opt.SetTxPoolExtra(n)
		return nil
	}
}
