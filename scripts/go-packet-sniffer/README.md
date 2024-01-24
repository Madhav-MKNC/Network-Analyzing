# Go Packet Sniffer

## Description
This Go Packet Sniffer is a simple tool for capturing and analyzing network traffic. It decodes various network protocols including Ethernet, IPv4, ICMP, TCP, and UDP, and can be used for educational purposes or network diagnostics.

## Requirements
- Go (version 1.15 or later)
- Administrative privileges for capturing packets
- Knowledge of the network interface you wish to capture packets from

## Installation
1. Clone the repository to your local machine.
2. Navigate to the directory containing the source code.

## Compilation
Run the following command to compile the sniffer:

```
go build sniffer.go
```

This will create an executable named 'sniffer' in the current directory.

## Running the Sniffer
To start packet capturing, run the compiled binary with administrative privileges. For example, on Linux, you would use:

```
sudo ./sniffer
```

Replace `./sniffer` with the path to your compiled binary if necessary.

## Important Notes
- The packet sniffer requires raw socket access, which typically requires administrative (root) privileges.
- Make sure you have the necessary permissions to capture network traffic in your network environment.
- The sniffer is configured for specific protocols and interfaces, and might require adjustments for your specific use case.
- This tool is for educational and diagnostic purposes. Be aware of privacy and legal considerations when capturing network traffic.

## Customization
You can modify the source code to capture different protocols or handle different network interfaces as per your requirements.
