package godhcp

import (
    "fmt"
    "net"
    "errors"
)

type Client struct {
    State       string
    HType 		HType
    CHAddr      net.HardwareAddr
    Hostname    string
    Options     OptionMap
}

func (c *Client) ProcessMessage(m *Message) (*Message, error) {
    //FIXME: Don't just overwrite our options, let's keep track of what we had before
    //  Options need some flag for which ones will be sent to server
    c.Options = m.Options
    
    switch m.Type {
    case DHCPDISCOVER:
        if c.State == "NEW" {
            c.State = "DHCPDISCOVER"
        }
        resp := c.PrepareOffer()
        return resp, nil
    }
    
    return nil, errors.New("Unknown Message Type Received")
}

func (c *Client) PrepareOffer() (*Message) {
    r := new(Message)
    r.Op = BOOTREPLY
    
    return r
}

func (c *Client) String() string {
    s := ""
    s += fmt.Sprintf("[%v] %v - %v", c.CHAddr, c.Hostname, c.State)
    return s
}