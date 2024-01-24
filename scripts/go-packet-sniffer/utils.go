package main

import (
    "fmt"
    "net"
    "strings"
)

// getMACAddr formats a MAC address in human-readable form
func getMACAddr(mac net.HardwareAddr) string {
    return mac.String()
}

// formatMultiLine formats multi-line data for cleaner output
func formatMultiLine(prefix string, data []byte, size int) string {
    if size <= len(prefix) {
        size = len(prefix) + 1
    }
    
    hexData := fmt.Sprintf("%x", data) // Convert to hex string
    wrapped := wrapText(hexData, size-len(prefix))

    var formattedLines []string
    for _, line := range wrapped {
        formattedLines = append(formattedLines, prefix+line)
    }

    return strings.Join(formattedLines, "\n")
}

// wrapText wraps a long string into lines of specified length
func wrapText(text string, lineLength int) []string {
    var lines []string
    for start := 0; start < len(text); start += lineLength {
        end := start + lineLength
        if end > len(text) {
            end = len(text)
        }
        lines = append(lines, text[start:end])
    }
    return lines
}
