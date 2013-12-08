package godhcp

import (
    "fmt"
    "net"
)

type Server struct {}

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
    }
    return nil
}

func (s *Server) Handle(m *Message) error {
    return nil
}

func (s *Server) Close() error {
    return nil
}