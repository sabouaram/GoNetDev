package SendRecv

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"github.com/sabouaram/GoNetDev/Protocols"
	"os/exec"
	"runtime"
	"strings"
)

type SendRecvInterface interface {
	SendFrame(frame []byte) (int, error)
	ReceiveFrame(byteSize int, chn chan Protocols.Frame) error
}

func NewSendRecvInterface() (SendRecvInterface, error) {
	nic, err := ListAndChooseInterface()
	if err != nil {
		return nil, err
	}
	return NewSendRecv(nic), nil
}

func isWindows() bool {
	return runtime.GOOS == "windows"
}

func isLinux() bool {
	return runtime.GOOS == "linux"
}

func ListAndChooseInterface() (string, error) {

	if runtime.GOOS == "windows" {
		devices, err := pcap.FindAllDevs()
		if err != nil {
			return "", err
		}

		fmt.Println("Available network interfaces:")
		for i, dev := range devices {
			fmt.Printf("%d: %s\n", i+1, dev.Name)
		}

		fmt.Print("Enter the number of the interface you want to use: ")
		var choice int
		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			return "", err
		}

		if choice < 1 || choice > len(devices) {
			return "", fmt.Errorf("Invalid choice")
		}

		selectedInterface := devices[choice-1].Name
		return selectedInterface, nil
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd := exec.Command("ifconfig")
		output, err := cmd.Output()
		if err != nil {
			return "", err
		}

		interfaces := []string{}
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, "flags=") {
				parts := strings.Split(line, ":")
				if len(parts) > 0 {
					interfaces = append(interfaces, parts[0])
				}
			}
		}

		// List available network interfaces
		fmt.Println("Available network interfaces:")
		for i, iface := range interfaces {
			fmt.Printf("%d: %s\n", i+1, iface)
		}

		// Ask the user to choose an interface
		fmt.Print("Enter the number of the interface you want to use: ")
		var choice int
		_, err = fmt.Scanf("%d", &choice)
		if err != nil {
			return "", err
		}

		if choice < 1 || choice > len(interfaces) {
			return "", fmt.Errorf("Invalid choice")
		}

		selectedInterface := interfaces[choice-1]
		return selectedInterface, nil
	} else {
		return "", fmt.Errorf("please install ifconfig on your OS")
	}
}

type WindowsSendRecv struct {
	Interface string
}

type UnixSendRecv struct {
	Interface string
}

func NewSendRecv(device string) SendRecvInterface {
	if isWindows() {
		return &WindowsSendRecv{
			Interface: device,
		}
	} else {
		return &UnixSendRecv{
			Interface: device,
		}
	}

	return nil
}
