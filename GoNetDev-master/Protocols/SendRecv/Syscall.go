package SendRecv

import (
	"github.com/sabouaram/GoNetDev/Protocols"
	"github.com/sabouaram/GoNetDev/Protocols/Utils"
	"github.com/sabouaram/GoNetDev/Protocols/Utils/GLogger"
	"net"
	"syscall"
)

func SyscallRawEth() (fd int) {
	fd, error := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, 0x0300)

	if error != nil {
		syscall.Close(fd)
		panic(error)
	}
	return fd
}

func ReceiveFrame(interface_name string, byte_size int, chn chan Protocols.Frame) error {
	fd := SyscallRawEth()
	err := syscall.BindToDevice(fd, interface_name)
	if err != nil {
		syscall.Close(fd)
		panic(err)
	}

	buffer := make([]byte, byte_size)
	logger := GLogger.GetInstance()
	logger.Printf("RX:%s", interface_name)
	Frames_received_metric := Utils.NewPromCounter("frames_received", "frames received per interface", []string{"interface_name"})
	Utils.Register(Frames_received_metric)
	for true {
		syscall.Recvfrom(fd, buffer, 0)
		Frame := Protocols.Parse(buffer)
		chn <-  Frame
		Frames_received_metric.With(Utils.NewLabel("interface_name", interface_name)).Inc()
		logger.Printf("RX:%s", interface_name)

	}
	return nil
}

func SendFrame(interface_name string, frame []byte) (count_bytes int, err error) {
	fd := SyscallRawEth()
	interface_info, err := net.InterfaceByName(interface_name)
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
	err = syscall.SetLsfPromisc(interface_name, true)
	if err != nil {
		return 0, err
	}
	n, err := syscall.Write(fd, frame)
	if err != nil {
		return 0, err
	}
	return n, err
}
