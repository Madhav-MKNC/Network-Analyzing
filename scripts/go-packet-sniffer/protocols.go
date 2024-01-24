package main

import (
    "encoding/binary"
    "net"
    "strings"
)

// Ethernet structure
type Ethernet struct {
    DestMAC   net.HardwareAddr
    SrcMAC    net.HardwareAddr
    ProtoType uint16
    Data      []byte
}

func NewEthernet(rawData []byte) *Ethernet {
    return &Ethernet{
        DestMAC:   net.HardwareAddr(rawData[0:6]),
        SrcMAC:    net.HardwareAddr(rawData[6:12]),
        ProtoType: binary.BigEndian.Uint16(rawData[12:14]),
        Data:      rawData[14:],
    }
}

// IPv4 structure
type IPv4 struct {
    Version        uint8
    HeaderLength   uint8
    TTL            uint8
    Protocol       uint8
    SourceIP       net.IP
    DestinationIP  net.IP
    Data           []byte
}

func NewIPv4(data []byte) *IPv4 {
    versionAndHeaderLength := data[0]
    return &IPv4{
        Version:        versionAndHeaderLength >> 4,
        HeaderLength:   (versionAndHeaderLength & 0x0F) * 4,
        TTL:            data[8],
        Protocol:       data[9],
        SourceIP:       net.IP(data[12:16]),
        DestinationIP:  net.IP(data[16:20]),
        Data:           data[20:],
    }
}

// ICMP structure
type ICMP struct {
    Type     uint8
    Code     uint8
    Checksum uint16
    Data     []byte
}

func NewICMP(data []byte) *ICMP {
    return &ICMP{
        Type:     data[0],
        Code:     data[1],
        Checksum: binary.BigEndian.Uint16(data[2:4]),
        Data:     data[4:],
    }
}

// TCP structure
type TCP struct {
    SrcPort      uint16
    DestPort     uint16
    Sequence     uint32
    Acknowledgment uint32
    DataOffset   uint8
    Flags        uint16
    WindowSize   uint16
    Checksum     uint16
    Urgent       uint16
    Data         []byte
}

func NewTCP(data []byte) *TCP {
    dataOffsetAndFlags := binary.BigEndian.Uint16(data[12:14])
    return &TCP{
        SrcPort:      binary.BigEndian.Uint16(data[0:2]),
        DestPort:     binary.BigEndian.Uint16(data[2:4]),
        Sequence:     binary.BigEndian.Uint32(data[4:8]),
        Acknowledgment: binary.BigEndian.Uint32(data[8:12]),
        DataOffset:   uint8(dataOffsetAndFlags >> 12),
        Flags:        dataOffsetAndFlags & 0x1FF, // Flags are 9 bits
        WindowSize:   binary.BigEndian.Uint16(data[14:16]),
        Checksum:     binary.BigEndian.Uint16(data[16:18]),
        Urgent:       binary.BigEndian.Uint16(data[18:20]),
        Data:         data[20:], // Assuming no options
    }
}

// UDP structure
type UDP struct {
    SrcPort      uint16
    DestPort     uint16
    Length       uint16
    Data         []byte
}

func NewUDP(data []byte) *UDP {
    return &UDP{
        SrcPort:      binary.BigEndian.Uint16(data[0:2]),
        DestPort:     binary.BigEndian.Uint16(data[2:4]),
        Length:       binary.BigEndian.Uint16(data[4:6]),
        Data:         data[8:],
    }
}

// HTTP structure
type HTTP struct {
    Data string
}

func NewHTTP(data []byte) *HTTP {
    return &HTTP{Data: string(data)}
}
