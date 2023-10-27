package Protocols

import (
	"encoding/binary"
	"fmt"
	"github.com/sabouaram/GoNetDev/Protocols/Utils"
	"math/bits"
	"reflect"
	"strconv"
	"strings"
)

type Ethernet struct {
	DestMacAddress   []byte
	SourceMacAddress []byte
	Dot1Q            []byte
	Type             []byte
}

func NewEthHeader() (Eth_header *Ethernet) {
	return &Ethernet{}
}
func (Header *Ethernet) BuildHeader(destmac string, src string, ntype uint16) {
	header := make([]byte, 18)
	Header.DestMacAddress = header[0:6]
	Header.SourceMacAddress = header[6:12]
	Header.Dot1Q = header[12:16]
	Header.Dot1Q = []byte{}
	Header.Type = header[16:18]

	if len(strings.Split(destmac, ":")) != 6 || len(strings.Split(src, ":")) != 6 {

	}

	hex_macs_string := append(strings.Split(destmac, ":"), strings.Split(src, ":")...)

	for i, v := range hex_macs_string {
		i_byte, err := strconv.ParseUint(v, 16, 8)
		if err == nil {
			binary.PutUvarint(header[i:], i_byte)
		}
	}

	binary.BigEndian.PutUint16(Header.Type, ntype)

}

func (Header *Ethernet) EthernetBytes() []byte {
	sRValue := reflect.ValueOf(Header).Elem()
	sRType := sRValue.Type()
	array := [][]byte{}
	for i := 0; i < sRType.NumField(); i++ {
		if len(sRValue.Field(i).Bytes()) != 0 {
			array = append(array, sRValue.Field(i).Bytes())
		}
	}
	return Utils.ConcatAppend(array)
}

func (Header *Ethernet) ParseEthernet(byte_slice []byte, isTrunked bool) {
	if isTrunked {
		Header.DestMacAddress = byte_slice[0:6]
		Header.SourceMacAddress = byte_slice[6:12]
		Header.Dot1Q = byte_slice[12:16]
		Header.Type = byte_slice[16:18]
	} else {
		Header.DestMacAddress = byte_slice[0:6]
		Header.SourceMacAddress = byte_slice[6:12]
		Header.Type = byte_slice[12:14]
	}
}

func (Header *Ethernet) TagDot1Q(VlanID int64, Priority int64) {
	TPID_s := fmt.Sprintf("%16b", 0x8100)
	PRI_s := fmt.Sprintf("%03b", Priority)
	CFI_s := fmt.Sprintf("%1b", 0)
	VID_s := fmt.Sprintf("%012b", VlanID)
	Tag, _ := strconv.ParseUint(TPID_s+PRI_s+CFI_s+VID_s, 2, 32)
	Tag_slice := make([]byte, 4)
	binary.BigEndian.PutUint32(Tag_slice, uint32(Tag))
	Header.Dot1Q = []byte{}
	Header.Dot1Q = append(Tag_slice)
}

func (Header *Ethernet) GetVlanID() (vlanid int64) {
	last2bytes := binary.BigEndian.Uint16(Header.Dot1Q[2:])
	vlanId, _ := strconv.ParseUint(fmt.Sprintf("%016b \n", bits.RotateLeft16(last2bytes, 4))[0:12], 2, 12)
	return int64(vlanId)
}

func (Header *Ethernet) GetPriority() (priority int64) {
	last2bytes := binary.BigEndian.Uint16(Header.Dot1Q[2:])
	Priority, _ := strconv.ParseUint(fmt.Sprintf("%03b \n", bits.RotateLeft16(last2bytes, 0))[0:3], 2, 3)
	return int64(Priority)
}
