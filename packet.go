package godhcp

import (
	"net"
    "fmt"
    "encoding/binary"
    "bytes"
)

/**
 Notes

Flags, 1st bit BROADCAST, rest MBZ
Options, 1st 4 bytes contain: 99, 130, 83, 99 magic cookie
RFC 1533 for options

*/
type OpCode byte

const (
	BOOTREQUEST OpCode = 1
	BOOTREPLY   OpCode = 2
)

func (t OpCode) String() string {
    switch t {
    case BOOTREQUEST:
        return "BOOTREQUEST"
    case BOOTREPLY:
        return "BOOTREPLY"
    }
    return "-INVALID-"
}

//See Assigned Numbers RFC
type HType byte

type Packet struct {
    Length      uint8 //length of Packet in Bytes
	Op			OpCode
	HType 		HType
	HLen		byte
	Hops 		uint8
	XId         []byte //4 bytes
	Secs 		uint16 //2 bytes
	Flags       []byte //2 bytes
	CIAddr      net.IP
	YIAddr      net.IP
	SIAddr      net.IP
	GIAddr      net.IP
	CHAddr      net.HardwareAddr
	SName       string //64 bytes
	File 		string //128 bytes
	Options 	[]Option	//Variable
}

func (p *Packet) FromBuffer(length int, b []byte) error {
    p.Length = uint8(length)
    p.Op = OpCode(b[0])
    p.HType = HType(b[1])
    p.HLen = b[2]
    p.Hops = uint8(b[3])
    p.XId = b[4:8]
    
    buf := bytes.NewReader(b[8:10])
    err := binary.Read(buf, binary.BigEndian, &p.Secs)
    if err != nil {
        return err
    }

    p.Flags = b[10:12]
    p.CIAddr = net.IP(b[12:16])
    p.YIAddr = net.IP(b[16:20])
    p.SIAddr = net.IP(b[20:24])
    p.GIAddr = net.IP(b[24:28])
    p.CHAddr = net.HardwareAddr(b[28:28 + p.HLen])
    p.SName = string(b[44:108])
    p.File = string(b[108:236])
    p.Options, err = ParseOptions(b[236:])
    return nil
}

func (p *Packet) ToBuffer() ([]byte, error) {
    buffer := make([]byte, p.Length)
    
    return buffer, nil
}

func (p *Packet) String() {
    s := fmt.Sprintln("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
    s += fmt.Sprintf("| %v ", p.Op.String())
    s += fmt.Sprintf("| %v ", p.HType)
    s += fmt.Sprintf("| %v ", p.HLen)
    s += fmt.Sprintf("| %v |\n", p.Hops)
    s += fmt.Sprintln("+---------------+---------------+---------------+---------------+")
    s += fmt.Sprintf("| %v |\n", p.XId)
    s += fmt.Sprintln("+-------------------------------+-------------------------------+")
    s += fmt.Sprintf("| %v s ", p.Secs)
    s += fmt.Sprintf("| %v |\n", p.Flags)
    s += fmt.Sprintln("+-------------------------------+-------------------------------+")
    s += fmt.Sprintln("| CI: "+ p.CIAddr.String())
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintln("| YI: "+ p.YIAddr.String())
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintln("| SI: "+ p.SIAddr.String())
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintln("| GI: "+ p.GIAddr.String())
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintf("| CH: %v |\n", p.CHAddr)
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintln("Server Name: ",p.SName)
    s += fmt.Sprintln("File Name: ",p.File)
    s += fmt.Sprintln("Options: ", p.Options)
    fmt.Print(s)
}

type DHCPMessageType byte

const (
	DHCPDISCOVER    DHCPMessageType = 1 // Broadcast C -> S
	DHCPOFFER       DHCPMessageType = 2 // S -> C, YIAddr set
	DHCPREQUEST     DHCPMessageType = 3 // C -> S
	DHCPDecline     DHCPMessageType = 4
	DHCPACK         DHCPMessageType = 5
	DHCPNAK         DHCPMessageType = 6
	DHCPRelease     DHCPMessageType = 7
	DHCPInform      DHCPMessageType = 8
)