# a simple packet capturing script

import socket

# Create a raw socket to sniff on
sniffer_socket = socket.socket(socket.AF_PACKET, socket.SOCK_RAW, socket.ntohs(3))

# Loop indefinitely to sniff packets
while True:
    # Receive a packet from the socket
    packet_data, _ = sniffer_socket.recvfrom(65536)

    # Extract the IP version from the packet
    version = packet_data[0] >> 4

    # If the packet is an IPv4 packet (version 4), extract the source and destination IP addresses
    if version == 4:
        # Extract the source and destination IP addresses from the IP header
        ip_header = packet_data[14:34]
        source_ip = socket.inet_ntoa(ip_header[12:16])
        dest_ip = socket.inet_ntoa(ip_header[16:20])

        # Print the source and destination IP addresses
        print(f"Source IP: {source_ip}, Destination IP: {dest_ip}")
