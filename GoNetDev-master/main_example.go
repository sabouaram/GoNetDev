package main

import (
	"github.com/sabouaram/GoNetDev/Protocols"
	"github.com/sabouaram/GoNetDev/Protocols/Const_Fields"
	"github.com/sabouaram/GoNetDev/Protocols/SendRecv"
)

func handleFrames(sendRecv SendRecv.SendRecvInterface) {
	chn := make(chan Protocols.Frame)
	go func() {
		err := sendRecv.ReceiveFrame(1024, chn)
		if err != nil {
			panic(err)
		}
	}()

	for s := range chn {
		// Example Replying to an ICMP Echo request
		if s.Icmph != nil && s.Icmph.GetType() == Const_Fields.ICMP_Type_Echo {
			s.Icmph.BuildICMPHeader(Const_Fields.ICMP_Type_Reply)
			s.Iph.ReverseSrc()
			_, err := sendRecv.SendFrame(s.FrameBytes()) // The interface name will be determined by the struct
			if err != nil {
				panic(err)
			}
		}
	}
}

func main() {

	sendRecv, err := SendRecv.NewSendRecvInterface()
	if err != nil {
		panic(err)
	}
	//frame := Protocols.NewFrame()
	//frame.Eth.BuildHeader("ff:ff:ff:ff:ff:ff", "08:00:27:dd:c1:1f", Const_Fields.Type_IPV4)
	//frame.Iph.BuildIPV4Header("192.168.10.10", "192.158.2.2", Const_Fields.ICMP_Type_Echo)
	//sendRecv.SendFrame(frame.FrameBytes())
	handleFrames(sendRecv)
}
