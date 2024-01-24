package main

import (
    "os"
    "time"
    "encoding/binary"
)

type Pcap struct {
    file *os.File
}

func NewPcap(filename string) *Pcap {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }

    var header = []byte{0xd4, 0xc3, 0xb2, 0xa1, 0x02, 0x00, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00}
    file.Write(header)

    return &Pcap{file: file}
}

func (p *Pcap) Write(data []byte) {
    ts := time.Now()
    tsSec := uint32(ts.Unix())
    tsUsec := uint32(ts.Nanosecond() / 1000)
    length := uint32(len(data))
    
    binary.Write(p.file, binary.LittleEndian, tsSec)
    binary.Write(p.file, binary.LittleEndian, tsUsec)
    binary.Write(p.file, binary.LittleEndian, length)
    binary.Write(p.file, binary.LittleEndian, length)
    p.file.Write(data)
}

func (p *Pcap) Close() {
    p.file.Close()
}
