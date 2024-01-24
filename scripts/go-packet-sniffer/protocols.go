package main

import (
    "encoding/binary"
    "net"
    "strings"
    "os"
    "time"
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

// Pcap structure
type Pcap struct {
    file *os.File
}

func NewPcap(filename string, linkType uint32) (*Pcap, error) {
    file, err := os.Create(filename)
    if err != nil {
        return nil, err
    }

    // Writing the Global Header for pcap
    globalHeader := []byte{
        0xd4, 0xc3, 0xb2, 0xa1, // Magic Number
        0x02, 0x00, 0x04, 0x00, // Version Major, Version Minor
        0x00, 0x00, 0x00, 0x00, // Thiszone
        0x00, 0x00, 0x00, 0x00, // Sigfigs
        0xff, 0xff, 0x00, 0x00, // Snaplen
        byte(linkType), byte(linkType >> 8), byte(linkType >> 16), byte(linkType >> 24), // Network
    }
    _, err = file.Write(globalHeader)
    if err != nil {
        return nil, err
    }

    return &Pcap{file: file}, nil
}

func (p *Pcap) Write(data []byte) error {
    ts := time.Now()
    tsSec := uint32(ts.Unix())
    tsUsec := uint32(ts.Nanosecond() / 1000)
    length := uint32(len(data))

    err := binary.Write(p.file, binary.LittleEndian, tsSec)
    if err != nil {
        return err
    }
    err = binary.Write(p.file, binary.LittleEndian, tsUsec)
    if err != nil {
        return err
    }
    err = binary.Write(p.file, binary.LittleEndian, length)
    if err != nil {
        return err
    }
    err = binary.Write(p.file, binary.LittleEndian, length)
    if err != nil {
        return err
    }
    _, err = p.file.Write(data)
    return err
}

func (p *Pcap) Close() {
    p.file.Close()
}
