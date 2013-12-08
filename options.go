package godhcp

//FIXME: Implement as interface so options can have separate implementations / checks
type Option struct {
    Code    OptionCode
    Length  int8
    Value   []byte
}

func (o *Option) String() string {
    return "Option"
}

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
    OPT_REQUESTEDIP                     OptionCode = 50
    OPT_LEASETIME                       OptionCode = 51
    OPT_OVERLOAD                        OptionCode = 52
    OPT_MESSAGETYPE                     OptionCode = 53
    OPT_SERVER                          OptionCode = 54
)

func ParseOptions(b []byte) ([]Option, error) {
    var opts []Option
    //cookie := b[0:4]
    
    //Read first byte, check against fixed length Options
    //If not fixed, check length and pass bytes into new Option
    return opts, nil
}