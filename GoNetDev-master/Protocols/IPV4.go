package Protocols

import (
	"encoding/binary"
	"fmt"
	"github.com/sabouaram/GoNetDev/Protocols/Const_Fields"
	"github.com/sabouaram/GoNetDev/Protocols/Utils"
	"math"
	"net"
	"strconv"
)

// RFC 791 INTERNET PROTOCOL VERSION 4
/* Developer: Salim BOU ARAM, e-mail: salimbouaram12@gmail.com */
type IP struct {
	Version_HL      []byte
	DS              []byte
	Total_Length    []byte
	Identification  []byte
	Flags           []byte
	TTL             []byte
	Protocol        []byte
	Header_checksum []byte
	Src_address     []byte
	Dst_address     []byte
}

type Packet struct {
	IPH  *IP
	Data []byte
}

func NewIpv4Header() (header *IP) {
	return &IP{}
}

func (Ip *IP) BuildIPV4Header(IPsrc string, IPdest string, Protocol uint8) {

	IPheader := []byte{Const_Fields.Version_Hl_IPV4, Const_Fields.DS, 0x00, 0x00, 0x00, 0x00, 0x40, 0x00, Const_Fields.TTL, Protocol, 0x00, 0x00}
	Ip.Version_HL = IPheader[0:1]
	Ip.DS = IPheader[1:2]
	Ip.Total_Length = IPheader[2:4]
	Ip.Identification = IPheader[4:6]
	Ip.Flags = IPheader[6:8]
	Ip.TTL = IPheader[8:9]
	Ip.Protocol = IPheader[9:10]
	Ip.Header_checksum = IPheader[10:12]
	Ip.Src_address = net.ParseIP(IPsrc).To4()
	Ip.Dst_address = net.ParseIP(IPdest).To4()
	Ip.HeaderChecksum()
}

func (Header *IP) IPV4Bytes() []byte {

	return Utils.ConcatAppend([][]byte{Header.Version_HL, Header.DS, Header.Total_Length, Header.Identification, Header.Flags, Header.TTL, Header.Protocol, Header.Header_checksum, Header.Src_address, Header.Dst_address})

}

func Fragment(data []byte, MTU int, Protocol uint8, IPsrc string, IPdest string) (Packets map[int]Packet) {
	results := make(map[int]Packet)
	data_size := len(data)
	fragment_count := int(math.Ceil(float64(data_size) / float64(MTU)))
	offset := 0
	Identification_bits := uint16(0x0000) // Initialize Identification_bits
	for k := 0; k < fragment_count; k++ {
		IP := NewIpv4Header()
		IP.BuildIPV4Header(IPsrc, IPdest, Protocol)
		IP.SetSRC(IPsrc)
		IP.SetDST(IPdest)
		Total_Length := 20
		var flags_hex uint16
		if k < fragment_count-1 {
			IP.Total_Length = []byte{uint8(uint16(Total_Length+MTU-34) >> 8), uint8(uint16(Total_Length+MTU-34) & 0xff)}
			IP.Identification = []byte{uint8(Identification_bits >> 8), uint8(Identification_bits & 0xff)}
			flags_hex = uint16(0x2000) // More fragments
			Identification_bits++
		} else {
			IP.Total_Length = []byte{uint8(uint16(Total_Length+data_size-offset) >> 8), uint8(uint16(Total_Length+data_size-offset) & 0xff)}
			IP.Identification = []byte{uint8(Identification_bits >> 8), uint8(Identification_bits & 0xff)}
			flags_hex = uint16(0x0000) // Last fragment
		}
		IP.Flags = []byte{uint8(flags_hex >> 8), uint8(flags_hex & 0xff)}
		IP.HeaderChecksum()
		fragmentSize := MTU - 34
		if k == fragment_count-1 {
			fragmentSize = data_size - offset
		}
		results[k] = Packet{IP, data[offset : offset+fragmentSize]}
		offset += fragmentSize
	}
	return results
}

func (Header *IP) HeaderChecksum() {
	intsum := int64(0)
	header_bytes := Header.IPV4Bytes()
	for i := 0; i <= 18; i += 2 {
		hex_string_2 := fmt.Sprintf("%x", header_bytes[i:i+2])
		uintsumf, _ := strconv.ParseUint(hex_string_2, 16, 16)
		intsum += int64(uintsumf)
		if intsum > 65536 {
			intsum -= 65536
		}
	}
	header_cheksum := uint16(intsum) ^ 0xffff
	binary.BigEndian.PutUint16(Header.Header_checksum, header_cheksum)
}

func (Header *IP) SetSRC(ipstring string) {
	Header.Src_address = []byte{}
	Header.Src_address = net.ParseIP(ipstring).To4()

}

func (Header *IP) SetDST(ipstring string) {
	Header.Dst_address = []byte{}
	Header.Dst_address = net.ParseIP(ipstring).To4()
}

func (Header *IP) ReverseSrc() {
	ipdst := Header.Dst_address
	Header.Dst_address = Header.Src_address
	Header.Src_address = ipdst
}

func (Header *IP) GetProtocol() uint8 {
	return Header.Protocol[0]
}

func (Header *IP) SetDS(label uint8) {
	Header.DS = []byte{}
	Header.DS = append(Header.DS, label)
}

func (Header *IP) ParseIPV4(byte_slice []byte) {
	Header.Version_HL = byte_slice[0:1]
	Header.DS = byte_slice[1:2]
	Header.Total_Length = byte_slice[2:4]
	Header.Identification = byte_slice[4:6]
	Header.Flags = byte_slice[6:8]
	Header.TTL = byte_slice[8:9]
	Header.Protocol = byte_slice[9:10]
	Header.Header_checksum = byte_slice[10:12]
	Header.Src_address = byte_slice[12:16]
	Header.Dst_address = byte_slice[16:20]

}
