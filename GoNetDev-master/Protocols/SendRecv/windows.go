//go:build windows
// +build windows

package SendRecv

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/sabouaram/GoNetDev/Protocols"
	"github.com/sabouaram/GoNetDev/Protocols/Utils"
	"github.com/sabouaram/GoNetDev/Protocols/Utils/GLogger"
)

func (w *WindowsSendRecv) SendFrame(frame []byte) (int, error) {
	handle, err := pcap.OpenLive(w.Interface, 65536, true, pcap.BlockForever)
	if err != nil {
		return 0, err
	}
	defer handle.Close()

	err = handle.WritePacketData(frame)
	if err != nil {
		return 0, err
	}

	return len(frame), nil

}
func (w *WindowsSendRecv) ReceiveFrame(byteSize int, chn chan Protocols.Frame) error {
	handle, err := pcap.OpenLive(w.Interface, int32(byteSize), true, pcap.BlockForever)
	if err != nil {
		return err
	}
	defer handle.Close()

	logger := GLogger.GetInstance()
	logger.Printf("RX:%s", w.Interface)
	FramesReceivedMetric := Utils.NewPromCounter("frames_received", "frames received per interface", []string{"interface_name"})
	Utils.Register(FramesReceivedMetric)

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		frame := Protocols.Parse(packet.Data())
		chn <- frame
		FramesReceivedMetric.With(Utils.NewLabel("interface_name", w.Interface)).Inc()
		logger.Printf("RX:%s", w.Interface)
	}

	return nil
}
