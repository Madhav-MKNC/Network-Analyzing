package main

import (
    "bytes"
    "encoding/binary"
    "fmt"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // Set up pcap file for capturing packets
    pcap, err := NewPcap("capture.pcap", 1) // linkType 1 for Ethernet
    if err != nil {
        log.Fatal(err)
    }
    defer pcap.Close()

    // Open a raw network interface
    iface, err := net.InterfaceByName("eth0") // Specify your interface name
    if err != nil {
        log.Fatal(err)
    }

    // Open a raw socket
    conn, err := net.ListenPacket("ip4:1", iface.Name)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    // Buffer to store packet data
    buf := make([]byte, 65535)

    // Handle exit signal
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        fmt.Println("Exiting...")
        os.Exit(0)
    }()

    for {
        n, _, err := conn.ReadFrom(buf)
        if err != nil {
            log.Fatal(err)
        }

        // Write raw data to pcap file
        if err := pcap.Write(buf[:n]); err != nil {
            log.Fatal(err)
        }

        // Process packet
        processPacket(buf[:n])
    }
}

func processPacket(packet []byte) {
    // Ethernet frame processing
    if len(packet) < 14 {
        return // Not enough data to process
    }
    eth := NewEthernet(packet)
    fmt.Printf("Ethernet Frame: Src: %s, Dst: %s, Type: %x\n", eth.SrcMAC, eth.DestMAC, eth.ProtoType)

    // Check for IPv4
    if eth.ProtoType == 0x0800 && len(eth.Data) >= 20 {
        ipv4 := NewIPv4(eth.Data)
        fmt.Printf("IPv4 Packet: Src: %s, Dst: %s, Protocol: %d\n", ipv4.SourceIP, ipv4.DestinationIP, ipv4.Protocol)

        switch ipv4.Protocol {
        case 1: // ICMP Protocol
            icmp := NewICMP(ipv4.Data)
            fmt.Printf("ICMP Packet: Type: %d, Code: %d\n", icmp.Type, icmp.Code)

        case 6: // TCP Protocol
            tcp := NewTCP(ipv4.Data)
            fmt.Printf("TCP Segment: Src Port: %d, Dst Port: %d\n", tcp.SrcPort, tcp.DestPort)

            // Check for HTTP (port 80)
            if tcp.SrcPort == 80 || tcp.DestPort == 80 {
                http := NewHTTP(tcp.Data)
                fmt.Printf("HTTP Data: %s\n", http.Data)
            }

        case 17: // UDP Protocol
            udp := NewUDP(ipv4.Data)
            fmt.Printf("UDP Segment: Src Port: %d, Dst Port: %d\n", udp.SrcPort, udp.DestPort)
        }
    }
}
