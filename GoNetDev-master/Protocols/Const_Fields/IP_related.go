package Const_Fields

const (
	Version_Hl_IPV4 = 0x45
	TTL = 0xff
	/*Classes (PHBs) COMBINED WITH ECN(2 bits set to 0) */
	DS = 0x00 //BEST EFFORT
    DS_ExpeditedForwarding= 0xb8
	DS_AF1 = 0x20
	DS_AF2 = 0x40 
	DS_AF3 = 0x60  
	DS_AF4 = 0x80  
	DS_NetworkControl = 0xc0
	/*Assuring Forwarding class 1 */
	DS_AF11 = 0x28  // Low drop prec
	DS_AF12 = 0x30  // Medium drop prec
	DS_AF13 = 0x38  // High drop prec
	/*Assuring Forwarding class 2 */
    DS_AF21 = 0x48  // Low drop prec
	DS_AF22 = 0x50  // Medium drop prec
	DS_AF23 = 0x58  // High drop prec
	/*Assuring Forwarding class 3 */
	DS_AF31 =  0x68 // Low drop prec
	DS_AF32 =  0x70 // Medium drop prec
	DS_AF33 =  0x78 // High drop prec
	/*Assuring Forwarding class 4 */
	DS_AF41 = 0x88  // Low drop prec
	DS_AF42 = 0x90  // Medium drop prec
	DS_AF43 = 0x98  // High drop prec




)
