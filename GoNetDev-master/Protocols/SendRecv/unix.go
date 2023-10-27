//go:build !windows
// +build !windows

package SendRecv

import (
	"github.com/sabouaram/GoNetDev/Protocols"
	"github.com/sabouaram/GoNetDev/Protocols/Utils"
	"github.com/sabouaram/GoNetDev/Protocols/Utils/GLogger"
	"net"
	"syscall"
)

func (u *UnixSendRecv) SendFrame(frame []byte) (int, error) {
	fd, error := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, 0x0300)
	if error != nil {
		syscall.Close(fd)
		panic(error)
	}
	interface_info, err := net.InterfaceByName(u.Interface)
	if err != nil {
		return 0, err
	}
	var haddr [8]byte
	copy(haddr[0:7], interface_info.HardwareAddr[0:7])
	addr := syscall.SockaddrLinklayer{
		Protocol: syscall.ETH_P_IP,
		Ifindex:  interface_info.Index,
		Halen:    uint8(len(interface_info.HardwareAddr)),
		Addr:     haddr,
	}
	err = syscall.Bind(fd, &addr)
	if err != nil {
		return 0, err
	}
	err = syscall.SetLsfPromisc(u.Interface, true)
	if err != nil {
		return 0, err
	}
	n, err := syscall.Write(fd, frame)
	if err != nil {
		return 0, err
	}
	return n, err
}
func (u *UnixSendRecv) ReceiveFrame(byteSize int, chn chan Protocols.Frame) error {
	fd, error := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, 0x0300)

	if error != nil {
		syscall.Close(fd)
		panic(error)
	}
	err := syscall.BindToDevice(fd, u.Interface)
	if err != nil {
		syscall.Close(fd)
		panic(err)
	}

	buffer := make([]byte, byteSize)
	logger := GLogger.GetInstance()
	logger.Printf("RX:%s", u.Interface)
	Frames_received_metric := Utils.NewPromCounter("frames_received", "frames received per interface", []string{"interface_name"})
	Utils.Register(Frames_received_metric)
	for true {
		syscall.Recvfrom(fd, buffer, 0)
		Frame := Protocols.Parse(buffer)
		chn <- Frame
		Frames_received_metric.With(Utils.NewLabel("interface_name", u.Interface)).Inc()
		logger.Printf("RX:%s", u.Interface)

	}
	return nil
}
