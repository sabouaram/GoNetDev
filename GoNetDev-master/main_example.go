package main

import (
	"github.com/sabouaram/GoNetDev/Protocols"
	"github.com/sabouaram/GoNetDev/Protocols/Const_Fields"
	"github.com/sabouaram/GoNetDev/Protocols/SendRecv"
)

func main() {

	/***********************************************************************************/
	// Frames handling and data processing (example replying to an ICMP echo request)
	chn := make(chan Protocols.Frame)
	go func() {
		err := SendRecv.ReceiveFrame("docker0", 1024, chn)
		if err != nil {
			panic(err)
		}
	}()

	for s := range chn {
		// Example Replying an ICMP Echo request
		if s.Icmph != nil && s.Icmph.GetType() == Const_Fields.ICMP_Type_Echo {
			s.Icmph.BuildICMPHeader(Const_Fields.ICMP_Type_Reply)
			s.Iph.ReverseSrc()
			_, err := SendRecv.SendFrame("enp0s3", s.FrameBytes())
			if err != nil {
				panic(err)
			}
		}
	}
	/************************************************************************************/
}
