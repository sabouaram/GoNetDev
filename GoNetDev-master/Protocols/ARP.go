package Protocols

import (
	"encoding/binary"
	"github.com/sabouaram/GoNetDev/Protocols/Const_Fields"
	"github.com/sabouaram/GoNetDev/Protocols/Utils"
	"net"
	"strconv"
	"strings"
)

type ARP struct {
	HardwareType     []byte
	ProtocolType     []byte
	HardwareSize     []byte
	ProtocolSize     []byte
	Operation        []byte
	SenderMACaddress []byte
	SenderIPaddress  []byte
	TargetMACaddress []byte
	TargetIPaddress  []byte
}

func NewArpHeader() *ARP {
	return &ARP{}
}

func (Arp *ARP) BuildARPHeader(HardwareType uint16, ProtocolType uint16, SenderIP string, TargetIP string, SenderMACaddress string, TargetMACaddress string, Operation uint16) {

	if HardwareType == Const_Fields.Hardware_type_Ethernet && ProtocolType == Const_Fields.Type_IPV4 {
		arp_header := []byte{0x00, 0x00, 0x00, 0x00, Const_Fields.ARP_ETH_HARDWARE_SIZE, Const_Fields.ARP_IPV4_PROTOCOL_SIZE, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		Arp.HardwareType = arp_header[0:2]
		Arp.ProtocolType = arp_header[2:4]
		Arp.HardwareSize = arp_header[4:5]
		Arp.ProtocolSize = arp_header[5:6]
		Arp.Operation = arp_header[6:8]
		Arp.SenderMACaddress = arp_header[8:14]
		Arp.SenderIPaddress = arp_header[14:18]
		Arp.TargetMACaddress = arp_header[18:24]
		Arp.TargetIPaddress = arp_header[24:28]
		binary.BigEndian.PutUint16(Arp.HardwareType, HardwareType)
		binary.BigEndian.PutUint16(Arp.ProtocolType, ProtocolType)
		IPsender := net.ParseIP(SenderIP).To4()
		IPtarget := net.ParseIP(TargetIP).To4()
		for i, v := range IPsender {
			Arp.SenderIPaddress[i] = v
		}
		for i, v := range IPtarget {
			Arp.TargetIPaddress[i] = v
		}

		for i, v := range strings.Split(SenderMACaddress, ":") {
			s_byte, _ := strconv.ParseUint(v, 16, 8)
			Arp.SenderMACaddress[i] = byte(s_byte)
		}

		if Operation == Const_Fields.ARP_Operation_request {
			binary.BigEndian.PutUint16(Arp.Operation, Operation)
			Arp.TargetMACaddress = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		} else if Operation == Const_Fields.ARP_Operation_reply {
			binary.BigEndian.PutUint16(Arp.Operation, Operation)
			for i, v := range strings.Split(SenderMACaddress, ":") {
				s_byte, _ := strconv.ParseUint(v, 16, 8)
				binary.PutUvarint(Arp.TargetMACaddress[i:], s_byte)
			}

		}
	}
}

func (arp *ARP) ARPBytes() []byte {
	return Utils.ConcatAppend([][]byte{arp.HardwareType, arp.ProtocolType, arp.HardwareSize, arp.ProtocolSize, arp.Operation, arp.SenderMACaddress, arp.SenderIPaddress, arp.TargetMACaddress, arp.TargetIPaddress})
}

func (arp *ARP) GetTargetMAC() []byte {
	return arp.TargetMACaddress
}

func (arp *ARP) SetTargetMAC(MAC []byte) {
	arp.TargetMACaddress = []byte{}
	arp.TargetMACaddress = MAC
}

func (arp *ARP) ParseARP(arp_byte_slice []byte) {
	arp.HardwareType = arp_byte_slice[0:2]
	arp.ProtocolType = arp_byte_slice[2:4]
	arp.HardwareSize = arp_byte_slice[4:5]
	arp.ProtocolSize = arp_byte_slice[5:6]
	arp.Operation = arp_byte_slice[6:8]
	arp.SenderMACaddress = arp_byte_slice[8:14]
	arp.SenderIPaddress = arp_byte_slice[14:18]
	arp.TargetMACaddress = arp_byte_slice[18:24]
	arp.TargetIPaddress = arp_byte_slice[24:28]
}
