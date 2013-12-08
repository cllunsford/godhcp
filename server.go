package godhcp

import (
    "fmt"
    "net"
    "errors"
    "bytes"
)

type Server struct {
    Clients     []*Client
}

func ListenAndServe() (error) {
    s := &Server{}
    conn, err := net.ListenUDP("udp4", &net.UDPAddr{
        IP: net.IPv4zero,
        Port: 67,
    })
    if err != nil {
        fmt.Println("server error:", err)
        return err
    }
    fmt.Println("listening on", conn.LocalAddr())
    
    defer conn.Close()
    
    return s.Serve(conn)
}

func (s *Server) Serve(conn *net.UDPConn) error {
    buffer := make([]byte, 1500) //FIXME: Find actual max packet size here
    
    for {
        length, remoteAddr, err := conn.ReadFrom(buffer)
        if err != nil {
            fmt.Println("server error:", err)
            return err
        }
        fmt.Println("Receive Packet", length, remoteAddr)
        
        m := &Message{}
        err = m.FromBuffer(length, buffer)
        if err != nil {
            fmt.Println("server error:", err)
            return err
        }
        fmt.Println(m.String())
        client, err := s.FindOrCreateClient(m)
        if err != nil {
            fmt.Println("Error creating client:", err)
        }
        
        resp, err := client.ProcessMessage(m)
        if err != nil {
            fmt.Println("Error parsing message:", err)
        }
        fmt.Println("Response: ", resp)
        
        fmt.Println("Clients:")
        for _, c := range s.Clients {
            fmt.Println(c.String())
        }
    }
    return nil
}

func (s *Server) Handle(m *Message) error {
    return nil
}

func (s *Server) Close() error {
    return nil
}

func (s *Server) FindOrCreateClient(m *Message) (*Client, error) {
    for _, c := range s.Clients {
        if bytes.Equal(c.CHAddr, m.CHAddr) {
            return c, nil
        }
    }
    
    if m.Type == DHCPDISCOVER {
        newc := &Client{}
        newc.State = "NEW"
        newc.CHAddr = m.CHAddr
        newc.HType = m.HType
        
        if v, ok := m.Options[OPT_HOSTNAME]; ok {
            newc.Hostname = string(v)
        }
        s.Clients = append(s.Clients, newc)
        return newc, nil
    }
    
    return nil, errors.New("New client without discovery")
}