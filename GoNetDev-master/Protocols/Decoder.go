package Protocols

import (
	"encoding/binary"
	"github.com/sabouaram/GoNetDev/Protocols/Const_Fields"
	"github.com/sabouaram/GoNetDev/Protocols/Utils/GLogger"
)

func Parse(byte_slice []byte) Frame {
	logger := GLogger.GetInstance()
	EthHeader := NewEthHeader()
	IPHeader := NewIpv4Header()
	ARPHeader := NewArpHeader()
	ICMPHeader := NewICMPHeader()
	padding := 0
	Tag1q := binary.BigEndian.Uint16(byte_slice[12:14])
	logger.Println(Tag1q)
	//In case of a Trunk link
	if Tag1q == 33024 {
		padding = 4
		EthHeader.ParseEthernet(byte_slice, true)
	} else {
		EthHeader.ParseEthernet(byte_slice, false)
	}
	logger.Printf("=>ETHERNET: SRC_MAC_ADDRESS:%x, DST_MAC_ADDRESS:%x \n", EthHeader.SourceMacAddress, EthHeader.DestMacAddress)
	switch binary.BigEndian.Uint16(EthHeader.Type) {

	case Const_Fields.Type_IPV4:
		{
			IPHeader.ParseIPV4(byte_slice[padding+14:])
			logger.Printf("  =>IPV4: SRC_IP_ADDRESS:%x, DST_IP_ADDRESS:%x \n", IPHeader.Src_address, IPHeader.Dst_address)
			switch uint8(IPHeader.Protocol[0]) {
			case Const_Fields.Type_ICMP:
				ICMPHeader.ParseICMP(byte_slice[padding+34:])
				logger.Println("    =>ICMP")
			case Const_Fields.Type_TCP:
				logger.Println("     =>TCP")
			case Const_Fields.Type_UDP:
				logger.Println("    =>UDP")
			}
		}
	case 0x0806:
		{
			ARPHeader.ParseARP(byte_slice[padding+14:])
			logger.Printf("  =>ARP: SRC_MAC_ADDRESS:%x, SRC_IP_ADDRESS:%x , TARGET_MAC_ADDRESS:%x, TARGET_IP_ADDRESS:%x, Operation: %x \n", ARPHeader.SenderMACaddress, ARPHeader.SenderIPaddress, ARPHeader.TargetMACaddress, ARPHeader.TargetIPaddress, ARPHeader.Operation)

		}
	}

	return Frame{EthHeader, ARPHeader, IPHeader, ICMPHeader}

}
