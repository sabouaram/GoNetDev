# GoNetDev

GoNetDev is an Ethernet based Golang Dataplane Development module designed simple for Linux/Unix systems.
#
Developer: Salim BOU ARAM.

## Installation & Compiling

```bash
go get -u  github.com/sabouaram/GoNetDev
cd $GOPATH/src/github.com/sabouaram/GoNetDev/ && go build
sudo ./main
```


## Usage
Ethernet Header
```go
package main
import ("github.com/sabouaram/GoNetDev/Protocols"
       	"github.com/sabouaram/GoNetDev/Protocols/Const_Fields"
        "fmt")
func main() {
frame := Protocols.NewFrame()
// BuildHeader func args: DSTMAC, SRCMAC, Protocol type
frame.Eth.BuildHeader("ff:ff:ff:ff:ff:ff", "08:00:27:dd:c1:1f", Const_Fields.Type_IPV4 )
// Optional: For IEEE 802.1q Tagging
frame.Eth.TagDot1Q(102, 6) //args: VlanID and 802.1p value
frame_bytes := frame.FrameBytes()
fmt.Printf("%x",frame_bytes)
}
```
ARP Request/Reply
```go
package main
import ("github.com/sabouaram/GoNetDev/Protocols"
       	"github.com/sabouaram/GoNetDev/Protocols/Const_Fields"
        "fmt")
func main(){
frame := Protocols.NewFrame()
frame.Eth.BuildHeader("ff:ff:ff:ff:ff:ff", "08:00:27:dd:c1:1f", Const_Fields.Type_ARP)
// ARP(Proto: IPV4 only supported) Func Args: srcIP, desIP, srcMAC, dstMAC, operation
// ARP REQUEST (for reply: last func arg: Const_Fields.ARP_Operation_reply) 
frame.Arph.BuildARPHeader(Const_Fields.Hardware_type_Ethernet, Const_Fields.Type_IPV4, "192.168.1.14", "192.168.1.222", "08:00:27:dd:c1:1f", "00:00:00:00:00:00",Const_Fields.ARP_Operation_request)
frame_bytes := frame.FrameBytes()
fmt.Printf("%x", frame_bytes)
}

```

ICMPv4 Echo/Reply

```go
package main
import ("github.com/sabouaram/GoNetDev/Protocols"
       	"github.com/sabouaram/GoNetDev/Protocols/Const_Fields"
         "fmt")
func main() {
frame := Protocols.NewFrame()
frame.Eth.BuildHeader("08:00:27:ff:23:22", "08:00:27:dd:c1:1f", Const_Fields.Type_IPV4)     
frame.Iph.BuildIPV4Header("192.168.1.14", "192.168.1.11", Const_Fields.Type_ICMP)
// BuildICMPHeader func args: ICMP Message type
frame.Icmph.BuildICMPHeader(Const_Fields.ICMP_Type_Reply)  
frame_bytes := frame.FrameBytes()
fmt.Printf("%x", frame_bytes)
}
```

IPv4 Fragmentation

```go
package main
import ("github.com/sabouaram/GoNetDev/Protocols"
       	"github.com/sabouaram/GoNetDev/Protocols/Const_Fields"
        "fmt")
func main() {
Data := make([]byte, 8900) // 8900 bytes of Data
frame := Protocols.NewFrame()
frame.Eth.BuildHeader("ff:ff:ff:ff:ff:ff", "08:00:27:dd:c1:1f", Const_Fields.Type_IPV4)
// Fragment func args: Data byte slice, MTU, IP Higher Protocol, IPSRC, IPDST
packets := Protocols.Fragment(Data, 1500, Const_Fields.Type_TCP, "192.168.0.12", "8.8.8.8")
for _,v := range packets 
  {
  frame.Iph = v
  frame_bytes := frame.FrameBytes()
  fmt.Printf("%x", frame_bytes)
  }
}
```

## Basic Ingress/Egress Processing (working on QoS feature and queuing management )
```go
package main
import ("github.com/sabouaram/GoNetDev/Protocols/SendRecv")
func main() {
frame := Protocols.NewFrame()
frame.Eth.BuildHeader("08:00:27:ff:23:22", "08:00:27:dd:c1:1f", Const_Fields.Type_IPV4)     
frame.Iph.BuildIPV4Header("192.168.1.14", "192.168.1.11", Const_Fields.Type_ICMP)
frame.Icmph.BuildICMPHeader(Const_Fields.ICMP_Type_Reply)
frame_bytes := frame.FrameBytes()

// Frames handling and data processing (example replying to an ICMP echo request)
  chn := make(chan Protocols.Frame)
	go func() {
		err := SendRecv.ReceiveFrame("enp0s3", 1024, chn)
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
}

```
## Adding a Protocol dynamically (working on it)
## Prometheus Metrics (working on it)
## also Tunneling / IPsec ...


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
