package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	target := "8.8.8.8"
	interval := 5 * time.Second

	for {
		addr, err := net.ResolveIPAddr("ip", target)
		if err != nil {
			fmt.Printf("Error resolving target IP address: %s\n", err.Error())
			continue
		}

		conn, err := net.DialIP("ip4:icmp", nil, addr)
		if err != nil {
			fmt.Printf("Error connecting to target IP address: %s\n", err.Error())
			continue
		}
		defer conn.Close()

		echo := make([]byte, 32)
		echo[0] = 8  // ICMP Echo Request
		echo[1] = 0  // Code
		echo[2] = 0  // Checksum (to be computed)
		echo[3] = 0  // Checksum (to be computed)
		echo[4] = 0  // Identifier (arbitrary value)
		echo[5] = 0  // Identifier (arbitrary value)
		echo[6] = 0  // Sequence number (arbitrary value)
		echo[7] = 0  // Sequence number (arbitrary value)

		// Compute ICMP checksum
		sum := uint32(0)
		for i := 0; i < len(echo); i += 2 {
			sum += uint32(echo[i+1])<<8 | uint32(echo[i])
		}
		sum = (sum >> 16) + (sum & 0xffff)
		sum += sum >> 16
		checksum := ^uint16(sum)

		// Set ICMP checksum in echo packet
		echo[2] = byte(checksum >> 8)
		echo[3] = byte(checksum)

		// Send ICMP echo request
		_, err = conn.Write(echo)
		if err != nil {
			fmt.Printf("Error sending ICMP echo request: %s\n", err.Error())
			continue
		}

		// Wait for ICMP echo reply
		reply := make([]byte, 32)
		_, err = conn.Read(reply)
		if err != nil {
			fmt.Printf("Error receiving ICMP echo reply: %s\n", err.Error())
			continue
		}

		// Print round-trip time
		rtt := time.Since(time.Now().Add(-interval)).Seconds() * 1000
		fmt.Printf("Ping to %s (%s) took %.3f ms\n", target, addr.String(), rtt)

		// Wait for next interval
		time.Sleep(interval)
	}
}
