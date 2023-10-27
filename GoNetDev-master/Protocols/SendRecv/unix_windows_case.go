//go:build windows
// +build windows

package SendRecv

import (
	"github.com/sabouaram/GoNetDev/Protocols"
)

func (u *UnixSendRecv) SendFrame(frame []byte) (int, error) {
	return 0, nil
}

func (u *UnixSendRecv) ReceiveFrame(byteSize int, chn chan Protocols.Frame) error {
	return nil
}
