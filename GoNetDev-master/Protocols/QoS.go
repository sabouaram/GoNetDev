package Protocols

import "github.com/sabouaram/GoNetDev/Protocols/Const_Fields"

//Layer2(Cos)/Layer3(dscp) mapping in L2L3QoSMap:
//================================
//802.1p: 0 1 2 3 4 5 6 7
//----------------------------
//dscp:   0 8 16 24 32 46 48 56 *

func QoSMapping() map[int64]uint8 {
	L2L3QoSMap := map[int64]uint8{
		0: Const_Fields.DS,                     //BEST EFFORT CLASS
		1: Const_Fields.DS_AF1,                 //AF1 CLASS
		2: Const_Fields.DS_AF2,                 //AF2 CLASS
		3: Const_Fields.DS_AF3,                 //AF3 CLASS
		4: Const_Fields.DS_AF4,                 //AF4 CLASS
		5: Const_Fields.DS_ExpeditedForwarding, //ExpeditedForwarding
		6: Const_Fields.DS_NetworkControl,      //Network control
	}
	return L2L3QoSMap
}

func (frame *Frame) LabelIPH(DS uint8) {
	frame.Iph.SetDS(DS)
}

// Function For marking
