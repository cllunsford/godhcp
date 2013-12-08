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

type Message struct {
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

func (m *Message) FromBuffer(length int, b []byte) error {
    m.Length = uint8(length)
    m.Op = OpCode(b[0])
    m.HType = HType(b[1])
    m.HLen = b[2]
    m.Hops = uint8(b[3])
    m.XId = b[4:8]
    
    buf := bytes.NewReader(b[8:10])
    err := binary.Read(buf, binary.BigEndian, &m.Secs)
    if err != nil {
        return err
    }

    m.Flags = b[10:12]
    m.CIAddr = net.IP(b[12:16])
    m.YIAddr = net.IP(b[16:20])
    m.SIAddr = net.IP(b[20:24])
    m.GIAddr = net.IP(b[24:28])
    m.CHAddr = net.HardwareAddr(b[28:28 + m.HLen])
    m.SName = string(b[44:108])
    m.File = string(b[108:236])
    m.Options, err = ParseOptions(b[236:])
    return nil
}

func (m *Message) ToBuffer() ([]byte, error) {
    buffer := make([]byte, m.Length)
    
    return buffer, nil
}

func (m *Message) String() string {
    s := fmt.Sprintln("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
    s += fmt.Sprintf("| %v ", m.Op.String())
    s += fmt.Sprintf("| %v ", m.HType)
    s += fmt.Sprintf("| %v ", m.HLen)
    s += fmt.Sprintf("| %v |\n", m.Hops)
    s += fmt.Sprintln("+---------------+---------------+---------------+---------------+")
    s += fmt.Sprintf("| %v |\n", m.XId)
    s += fmt.Sprintln("+-------------------------------+-------------------------------+")
    s += fmt.Sprintf("| %v s ", m.Secs)
    s += fmt.Sprintf("| %v |\n", m.Flags)
    s += fmt.Sprintln("+-------------------------------+-------------------------------+")
    s += fmt.Sprintln("| CI: "+ m.CIAddr.String())
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintln("| YI: "+ m.YIAddr.String())
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintln("| SI: "+ m.SIAddr.String())
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintln("| GI: "+ m.GIAddr.String())
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintf("| CH: %v |\n", m.CHAddr)
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintln("| Server Name: ",m.SName)
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintln("| File Name: ",m.File)
    s += fmt.Sprintln("+---------------------------------------------------------------+")
    s += fmt.Sprintln("Options: ", m.Options)
    return s
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