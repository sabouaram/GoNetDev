package Protocols

import (
	"encoding/binary"
	"fmt"
	"github.com/sabouaram/GoNetDev/Protocols/Const_Fields"
	"github.com/sabouaram/GoNetDev/Protocols/Utils"
	"reflect"
	"strconv"
)

// RFC 792 INTERNET CONTROL MESSAGE PROTOCOL
/* Developer: Salim BOU ARAM, e-mail: salimbouaram12@gmail.com */

type ICMP struct {
	Type       []byte
	Code       []byte
	Checksum   []byte
	Identifier []byte
	SequenceN  []byte
	Payload    []byte
}

func NewICMPHeader() (header *ICMP) {
	return &ICMP{}
}
func (ICMPH *ICMP) BuildICMPHeader(Type uint8) {
	var Header []byte
	switch Type {
	case Const_Fields.ICMP_Type_Echo:
		{
			Header = make([]byte, 40)
			ICMPH.Type = Header[0:1]
			ICMPH.Code = Header[1:2]
			ICMPH.Checksum = Header[2:4]
			ICMPH.Identifier = Header[4:6]
			ICMPH.SequenceN = Header[6:8]
			ICMPH.Payload = Header[8:]
			ICMPH.Type = []byte{Type}
			ICMPH.Code = []byte{0x00}
			ICMPH.Checksum = []byte{0x00, 0x00}
			ICMPH.Identifier = []byte{0x00, 0x01}
			ICMPH.SequenceN = []byte{0x00, 0x4f}
			ICMPH.Payload = []byte{0x74, 0x65, 0x73, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
			ICMPH.CheckSum()
		}

	case Const_Fields.ICMP_Type_Reply:
		{
			Header = make([]byte, 40)
			ICMPH.Type = Header[0:1]
			ICMPH.Code = Header[1:2]
			ICMPH.Checksum = Header[2:4]
			ICMPH.Identifier = Header[4:6]
			ICMPH.SequenceN = Header[6:8]
			ICMPH.Payload = Header[8:]
			ICMPH.Type = []byte{Type}
			ICMPH.Code = []byte{0x00}
			ICMPH.Checksum = []byte{0x00, 0x00}
			ICMPH.Identifier = []byte{0x00, 0x001}
			ICMPH.SequenceN = []byte{0x00, 0x54}
			ICMPH.Payload = []byte{0x74, 0x65, 0x73, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
			ICMPH.CheckSum()
		}

	case Const_Fields.ICMP_Type_Exceeded:
		{
			// -> TO DO
		}

	case Const_Fields.ICMP_Type_Unreachable:
		{
			// -> TO DO
		}

	case Const_Fields.ICMP_Type_Paremeter_problem:
		{
			// -> TO DO
		}
	case Const_Fields.ICMP_Type_Source_quench:
		{
			// -> TO DO
		}

	case Const_Fields.ICMP_Type_Redirect:
		{
			// -> TO DO
		}
	case Const_Fields.ICMP_Type_Timestamp:
		{
			// -> TO DO
		}
	case Const_Fields.ICMP_Type_TimestampReply:
		{
			// -> TO DO
		}
	case Const_Fields.ICMP_Type_InformationRequest:
		{
			// -> TO DO
		}
	case Const_Fields.ICMP_Type_InformationReply:
		{
			// -> TO DO
		}

	}

}

func (ICMPH *ICMP) CheckSum() {
	byte_array := ICMPH.ICMPBytes()
	var intsum int64 = 0
	for i := 0; i <= len(byte_array)-1; i += 2 {
		uintsumf, _ := strconv.ParseUint(fmt.Sprintf("%x", byte_array[i:i+2]), 16, 16)
		intsum += int64(uintsumf)
		if intsum > 65536 {
			intsum -= 65536
		}
	}
	cheksum := uint16(intsum) ^ 0xffff
	binary.BigEndian.PutUint16(ICMPH.Checksum, cheksum)
}

func (ICMPH *ICMP) SetPayload(data []byte) {
	ICMPH.Payload = []byte{}
	ICMPH.Payload = data
	ICMPH.CheckSum()
}

func (ICMPH *ICMP) GetIdentifier() (Identifier []byte) {
	return ICMPH.Identifier
}

func (ICMPH *ICMP) SetIdentifier(Identifier []byte) {
	ICMPH.Identifier = []byte{}
	ICMPH.Identifier = Identifier
	ICMPH.CheckSum()
}

func (ICMPH *ICMP) GetSequenceN() (SequenceN []byte) {
	return ICMPH.SequenceN
}

func (ICMPH *ICMP) SetSequenceN(SequenceN []byte) {
	ICMPH.SequenceN = []byte{}
	ICMPH.SequenceN = SequenceN
	ICMPH.CheckSum()

}

func (ICMPH *ICMP) GetType() uint8 {
	array := [8]uint8{}
	copy(array[:], ICMPH.Type)
	return array[0]
}

func (ICMPH *ICMP) ParseICMP(byte_slice []byte) {
	ICMPH.Type = byte_slice[0:1]
	ICMPH.Code = byte_slice[1:2]
	ICMPH.Checksum = byte_slice[2:4]
	ICMPH.Identifier = byte_slice[4:6]
	ICMPH.SequenceN = byte_slice[6:8]
	ICMPH.Payload = byte_slice[8:40]

}

func (ICMPH *ICMP) ICMPBytes() []byte {
	sRValue := reflect.ValueOf(ICMPH).Elem()
	sRType := sRValue.Type()
	array := [][]byte{}
	for i := 0; i < sRType.NumField(); i++ {
		if len(sRValue.Field(i).Bytes()) != 0 {
			array = append(array, sRValue.Field(i).Bytes())
		}
	}

	return Utils.ConcatAppend(array)

}
