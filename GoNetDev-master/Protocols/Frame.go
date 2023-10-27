package Protocols

import (
	"github.com/sabouaram/GoNetDev/Protocols/Utils"
	"reflect"
)

var Buffer_Channel chan Frame

type Frame struct {
	Eth   *Ethernet
	Arph  *ARP
	Iph   *IP
	Icmph *ICMP
	/*others*/
}

func NewFrame() *Frame {
	frame := Frame{}
	frame.Eth = NewEthHeader()
	frame.Arph = NewArpHeader()
	frame.Iph = NewIpv4Header()
	frame.Icmph = NewICMPHeader()
	return &frame
}

func (frame *Frame) FrameBytes() (byte_array []byte) {
	sRValue := reflect.ValueOf(frame).Elem()
	sRType := sRValue.Type()
	array := [][]byte{}
	for i := 0; i < sRType.NumField(); i++ {
		value := sRValue.Field(i).Interface()
		switch sRType.Field(i).Name {
		case "Eth":
			{
				eth, _ := value.(*Ethernet)
				array = append(array, eth.EthernetBytes())
			}
		case "Arph":
			{
				arph, _ := value.(*ARP)
				array = append(array, arph.ARPBytes())
			}
		case "Iph":
			{
				iph, _ := value.(*IP)
				array = append(array, iph.IPV4Bytes())
			}

		case "Icmph":
			{
				icmph, _ := value.(*ICMP)
				array = append(array, icmph.ICMPBytes())
			}

		}
	}
	return Utils.ConcatAppend(array)

}
