//go:build !windows
// +build !windows

package SendRecv

import "github.com/sabouaram/GoNetDev/Protocols"

func (u *WindowsSendRecv) SendFrame(frame []byte) (int, error) {
	return 0, nil
}

func (u *WindowsSendRecv) ReceiveFrame(byteSize int, chn chan Protocols.Frame) error {
	return nil
}
