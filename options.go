package godhcp

import (
    "bytes"
)

type OptionCode byte

const(
    // RFC 1497 Vendor Extensions
    OPT_PAD                             OptionCode = 0
    OPT_END                             OptionCode = 255
    OPT_SUBNETMASK                      OptionCode = 1
    OPT_TIMEOFFSET                      OptionCode = 2 
    OPT_ROUTER                          OptionCode = 3  
    OPT_TIMESERVER                      OptionCode = 4  
    OPT_NAMESERVER                      OptionCode = 5  
    OPT_DOMAINNAMESERVER                OptionCode = 6
    OPT_LOGSERVER                       OptionCode = 7
    OPT_COOKIESERVER                    OptionCode = 8
    OPT_LPRSERVER                       OptionCode = 9
    OPT_IMPRESSERVER                    OptionCode = 10
    OPT_RESOURCELOCATIONSERVER          OptionCode = 11
    OPT_HOSTNAME                        OptionCode = 12
    OPT_BOOTFILESIZE                    OptionCode = 13
    OPT_MERITDUMPFILE                   OptionCode = 14
    OPT_DOMAINNAME                      OptionCode = 15
    OPT_SWAPSERVER                      OptionCode = 16
    OPT_ROOTPATH                        OptionCode = 17
    OPT_EXTENSIONSPATH                  OptionCode = 18
    // IP Layer Parameters per Host
    OPT_IPFORWARDING                    OptionCode = 19
    OPT_NONLOCALSOURCEROUTE             OptionCode = 20
    OPT_POLICYFILTER                    OptionCode = 21
    OPT_MAXIMUMDATAGRAMREASSEMBLYSIZE   OptionCode = 22
    OPT_DEFAULTIPTTL                    OptionCode = 23
    OPT_PATHMTUAGINGTIMEOUT             OptionCode = 24
    // IP Layer Parameters per Interface
    // Link Layer Parameters per Interface
    // TCP Parameters
    // Application and Service Parameters
    // DHCP Extensions
    OPT_REQUESTED_IP                    OptionCode = 50
    OPT_LEASE_TIME                      OptionCode = 51
    OPT_OVERLOAD                        OptionCode = 52
    OPT_MESSAGE_TYPE                    OptionCode = 53
    OPT_SERVER                          OptionCode = 54
    OPT_PARAMETER_REQUEST_LIST          OptionCode = 55
)

func (t OptionCode) String() string {
    switch t {
    case OPT_HOSTNAME:
        return "OPT_HOSTNAME"
    case OPT_MESSAGE_TYPE:
        return "OPT_MESSAGE_TYPE"
    case OPT_PARAMETER_REQUEST_LIST:
        return "OPT_PARAMETER_REQUEST_LIST"
    }
    return "-UNKNOWN-"
}

type OptionMap map[OptionCode][]byte


//TODO: this really needs a clean up / simplify
func ParseOptions(b []byte) (OptionMap, error) {
    opts := make(OptionMap)
    opBuf := make([]byte, 1)
    
    //cookie := b[0:4]
    rdr := bytes.NewReader(b[4:])
    
    for {
        _, err := rdr.Read(opBuf)
        nextOpCode := OptionCode(opBuf[0])
        if nextOpCode == OPT_END {
            return opts, nil
        }
        if err != nil {
            return opts, err
        }
        
        sizeBuf := make([]byte, 1)
        _, err = rdr.Read(sizeBuf)
        if err != nil {
            return opts, err
        }
        size := int8(sizeBuf[0])
        
        value := make([]byte, size)
        _, err = rdr.Read(value)
        if err != nil {
            return opts, err
        }
        opts[nextOpCode] = value
    }
    
    //Read first byte, check against fixed length Options
    //If not fixed, check length and pass bytes into new Option
    return opts, nil
}

type OptionGeneric struct {
    Code    OptionCode
    Length  int8
    Value   []byte
}

func (o *OptionGeneric) String() string {
    return "Generic Option"
}

type OptionMessageType struct {
    Code    OptionCode
    Length  int8
    Value   []byte
}

func (o *OptionMessageType) String() string {
    return "Option Message Type"
}