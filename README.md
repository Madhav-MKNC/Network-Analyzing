# Network-Analyzing


<a href="https://github.com/Madhav-MKNC/Python-Packet-Sniffer">More Here</a>


<!-- all merged here -->



<!-- NETWORK DATA UNITS -->
# Overview of Network Data Units
<p></p>


<p>In computer networking, data is transmitted in discrete units called <strong>Network Data Units (NDUs)</strong>. These units help to ensure that data is transmitted efficiently and reliably across the network. In this overview, we will discuss the different types of NDUs and their characteristics.</p>
<h2>Types of Network Data Units</h2>
<h3>1. Bit</h3>
<p>The smallest unit of data in a computer network is a <strong>bit</strong>. It can have a value of either 0 or 1.</p>
<h3>2. Byte</h3>
<p>A <strong>byte</strong> is a unit of data that consists of 8 bits. It is the basic building block of most computer systems and is used to represent characters, numbers, and other types of data.</p>
<h3>3. Packet</h3>
<p>A <strong>packet</strong> is a unit of data that is transmitted across a network. It consists of a header and a payload. The header contains information about the packet, such as its source and destination addresses, while the payload contains the actual data being transmitted.</p>
<h3>4. Frame</h3>
<p>A <strong>frame</strong> is a unit of data that is used in local area networks (LANs). It consists of a header, data, and a trailer. The header and trailer contain information about the frame, such as its source and destination addresses, while the data contains the actual information being transmitted.</p>
<h3>5. Segment</h3>
<p>A <strong>segment</strong> is a unit of data that is used in transport layer protocols, such as TCP (Transmission Control Protocol). It consists of a header and data. The header contains information about the segment, such as its source and destination ports, while the data contains the actual information being transmitted.</p>
<h2>Comparison of Network Data Units</h2>
<p>The following table summarizes the characteristics of the different types of NDUs:</p>
<table>
    <thead>
        <tr>
            <th>NDU</th>
            <th>Size</th>
            <th>Used in</th>
            <th>Example</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Bit</td>
            <td>1</td>
            <td>All layers</td>
            <td>0 or 1</td>
        </tr>
        <tr>
            <td>Byte</td>
            <td>8 bits</td>
            <td>All layers</td>
            <td>'a', 27</td>
        </tr>
        <tr>
            <td>Packet</td>
            <td>Variable</td>
            <td>Network layer</td>
            <td>IP packet</td>
        </tr>
        <tr>
            <td>Frame</td>
            <td>Variable</td>
            <td>Data link layer</td>
            <td>Ethernet frame</td>
        </tr>
        <tr>
            <td>Segment</td>
            <td>Variable</td>
            <td>Transport layer (TCP)</td>
            <td>TCP segment</td>
        </tr>
    </tbody>
</table>
<h2>Conclusion</h2>
<p>Network Data Units play a crucial role in ensuring the efficient and reliable transmission of data across computer networks. By understanding the different types of NDUs and their characteristics, network engineers can design and optimize their networks for maximum performance and reliability.</p>





<!-- NETWORK PROTOCOL DATA UNITS -->
# Network Protocol Data Units (NPDUs)
<p></p>

<h1>Overview of Network Data Units</h1>
<p>In computer networking, data is transmitted in discrete units called <strong>Network Data Units (NDUs)</strong>. These units help to ensure that data is transmitted efficiently and reliably across the network. In this overview, we will discuss the different types of NDUs and their characteristics for Ethernet frames, WiFi frames, IP datagrams, TCP segments, and UDP.</p>
<h2>Ethernet Frames</h2>
<p>Ethernet frames are used in Ethernet networks to transmit data between devices. They consist of a header, data, and a
    trailer. The header contains information such as the source and destination MAC addresses, while the trailer
    contains error detection information.</p>
<table>
    <thead>
        <tr>
            <th>NDU</th>
            <th>Size</th>
            <th>Used in</th>
            <th>Example</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Ethernet Frame</td>
            <td>64-1518 bytes</td>
            <td>Data link layer</td>
            <td>Ethernet II, IEEE 802.3 frames</td>
        </tr>
    </tbody>
</table>
<h2>WiFi Frames</h2>
<p>WiFi frames are used in wireless networks to transmit data between devices. They consist of a header, data, and a trailer. The header contains information such as the source and destination MAC addresses, while the trailer contains error detection information.</p>
<table>
    <thead>
        <tr>
            <th>NDU</th>
            <th>Size</th>
            <th>Used in</th>
            <th>Example</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>WiFi Frame</td>
            <td>Variable</td>
            <td>Data link layer</td>
            <td>IEEE 802.11 frames</td>
        </tr>
    </tbody>
</table>
<h2>IP Datagrams</h2>
<p>IP datagrams are used to transmit data across IP networks. They consist of a header and data. The header contains information such as the source and destination IP addresses, while the data contains the actual information being transmitted.</p>
<table>
    <thead>
        <tr>
            <th>NDU</th>
            <th>Size</th>
            <th>Used in</th>
            <th>Example</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>IP Datagram</td>
            <td>Variable</td>
            <td>Network layer</td>
            <td>IPv4, IPv6 datagrams</td>
        </tr>
    </tbody>
</table>
<h2>TCP Segments</h2>
<p>TCP segments are used in transport layer protocols, such as TCP, to transmit data across networks. They consist of a header and data. The header contains information such as the source and destination port numbers, while the data contains the actual information being transmitted.</p>
<table>
    <thead>
        <tr>
            <th>NDU</th>
            <th>Size</th>
            <th>Used in</th>
            <th>Example</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>TCP Segment</td>
            <td>Variable</td>
            <td>Transport layer (TCP)</td>
            <td>TCP segments</td>
        </tr>
    </tbody>
</table>
<h2>UDP Datagrams</h2>
<p>UDP datagrams are used in transport layer protocols, such as UDP, to transmit data across networks. They consist of a header and data. The header contains information such as the source and destination port numbers, while the data contains the actual information being transmitted.</p>
<table>
    <thead>
        <tr>
            <th>NDU</th>
            <th>Size</th>
            <th>Used in</th>
            <th>Example</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>UDP Datagram</td>
            <td>Variable</td>
            <td>Transport layer (UDP)</td>
            <td>UDP datagrams</td>
        </tr>
    </tbody>
</table>
<h2>Conclusion</h2>
<p>By understanding the different types of NDUs used in computer networking, network engineers can design and optimize their networks for maximum performance and reliability. Ethernet frames, WiFi frames, IP datagrams, TCP segments, and UDP datagrams all play an important role in the efficient and reliable transmission of data across computer networks.</p>




<!-- MORE ON THESE FRAMES -->
# STRCUTURE AND FORMAT OF THESE FRAMES
<p></p>

<h2>Ethernet Frames</h2>
<p>Ethernet frames consist of a header, data, and trailer.</p>
<h3>Ethernet Frame Header</h3>
<table>
    <thead>
        <tr>
            <th>Field</th>
            <th>Size (bits)</th>
            <th>Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Preamble</td>
            <td>7 bytes (56 bits)</td>
            <td>A sequence of alternating 0s and 1s to allow for synchronization between devices</td>
        </tr>
        <tr>
            <td>Start Frame Delimiter</td>
            <td>1 byte (8 bits)</td>
            <td>Marks the beginning of the frame</td>
        </tr>
        <tr>
            <td>Destination MAC Address</td>
            <td>6 bytes (48 bits)</td>
            <td>Specifies the destination device</td>
        </tr>
        <tr>
            <td>Source MAC Address</td>
            <td>6 bytes (48 bits)</td>
            <td>Specifies the sending device</td>
        </tr>
        <tr>
            <td>EtherType/Length</td>
            <td>2 bytes (16 bits)</td>
            <td>Specifies the type of data being carried in the frame or the length of the frame</td>
        </tr>
        <tr>
            <td>VLAN Tag (Optional)</td>
            <td>4 bytes (32 bits)</td>
            <td>Specifies a VLAN tag, if present</td>
        </tr>
    </tbody>
</table>
<h3>Ethernet Frame Data</h3>
<p>The data in an Ethernet frame can vary in size from 46 to 1500 bytes.</p>
<h3>Ethernet Frame Trailer</h3>
<p>The trailer contains an error detection mechanism called the Frame Check Sequence (FCS).</p>
<table>
    <thead>
        <tr>
            <th>Field</th>
            <th>Size (bits)</th>
            <th>Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>FCS</td>
            <td>4 bytes (32 bits)</td>
            <td>A CRC-32 checksum of the entire frame, used for error detection</td>
        </tr>
    </tbody>
</table>
<h2>WiFi Frames</h2>
<p>WiFi frames consist of a header, data, and trailer.</p>
<h3>WiFi Frame Header</h3>
<table>
    <thead>
        <tr>
            <th>Field</th>
            <th>Size (bits)</th>
            <th>Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Frame Control</td>
            <td>2 bytes (16 bits)</td>
            <td>Contains various control information, such as the frame type and subtype</td>
        </tr>
        <tr>
            <td>Duration/ID</td>
            <td>2 bytes (16 bits)</td>
            <td>Specifies the duration of the transmission or the identifier of the frame</td>
        </tr>
        <tr>
            <td>Destination MAC Address</td>
            <td>6 bytes (48 bits)</td>
            <td>Specifies the destination device</td>
        </tr>
        <tr>
            <td>Source MAC Address</td>
            <td>6 bytes (48 bits)</td>
            <td>Specifies the sending device</td>
        </tr>
        <tr>
            <td>BSSID (Basic Service Set Identifier)</td>
            <td>6 bytes (48 bits)</td>
            <td>Specifies the MAC address of the access point</td>
        </tr>
        <tr>
            <td>Sequence Control</td>
            <td>2 bytes (16 bits)</td>
            <td>Used to provide sequence numbers for fragmentation and reassembly</td>
        </tr>
        <tr>
            <td>QoS Control (Optional)</td>
            <td>2 bytes (16 bits)</td>
            <td>Specifies Quality of Service information, if present</td>
        </tr>
        <tr>
            <td>HT Control (Optional)</td>
            <td>4 bytes (32 bits)</td>
            <td>Specifies High Throughput control information, if present</td>
        </tr>
    </tbody>
</table>
<h3>WiFi Frame Data</h3>
<p>The data in a WiFi frame can vary in size from 0 to 2312 bytes.</p>
<h3>WiFi Frame Trailer</h3>
<p>The trailer contains an error detection mechanism called the Frame Check Sequence (FCS).</p>
<table>
    <thead>
        <tr>
            <th>Field</th>
            <th>Size (bits)</th>
            <th>Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>FCS</td>
            <td>4 bytes (32 bits)</td>
            <td>A CRC-32 checksum of the entire frame, used for error detection</td>
        </tr>
    </tbody>
</table>
<h2>IP Datagrams</h2>
<p>IP datagrams consist of a header and data.</p>
<h3>IP Datagram Header</h3>
<table>
    <thead>
        <tr>
            <th>Field</th>
            <th>Size (bits)</th>
            <th>Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Version</td>
            <td>4 bits</td>
            <td>Specifies the IP protocol version</td>
        </tr>
        <tr>
            <td>Header Length</td>
            <td>4 bits</td>
            <td>Specifies the length of the header in 32-bit words</td>
        </tr>
        <tr>
            <td>Type of Service</td>
            <td>8 bits</td>
            <td>Specifies the type of service, such as priority or reliability</td>
        </tr>
        <tr>
            <td>Total Length</td>
            <td>16 bits</td>
            <td>Specifies the total length of the datagram, including the header and data</td>
        </tr>
        <tr>
            <td>Identification</td>
            <td>16 bits</td>
            <td>Used for fragmentation and reassembly</td>
        </tr>
        <tr>
            <td>Flags</td>
            <td>3 bits</td>
            <td>Used for fragmentation and reassembly</td>
        </tr>
        <tr>
            <td>Fragment Offset</td>
            <td>13 bits</td>
            <td>Used for fragmentation and reassembly</td>
        </tr>
        <tr>
            <td>Time to Live</td>
            <td>8 bits</td>
            <td>Specifies the maximum number of hops the datagram can take before being discarded</td>
        </tr>
        <tr>
            <td>Protocol</td>
            <td>8 bits</td>
            <td>Specifies the protocol used in the data portion of the datagram</td>
        </tr>
        <tr>
            <td>Header Checksum</td>
            <td>16 bits</td>
            <td>A checksum of the header used for error detection</td>
        </tr>
        <tr>
            <td>Source IP Address</td>
            <td>32 bits</td>
            <td>Specifies the IP address of the sending device</td>
        </tr>
        <tr>
            <td>Destination IP Address</td>
            <td>32 bits</td>
            <td>Specifies the IP address of the destination device</td>
        </tr>
    </tbody>
</table>
<h3>IP Datagram Data</h3>
<p>The data in an IP datagram can vary in size up to 65,535 bytes.</p>
<h2>TCP Segments</h2>
<p>TCP segments consist of a header, data, and optional trailer.</p>
<h3>TCP Segment Header</h3>
<table>
    <thead>
        <tr>
            <th>Field</th>
            <th>Size (bits)</th>
            <th>Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Source Port</td>
            <td>16 bits</td>
            <td>Specifies the port number of the sending process</td>
        </tr>
        <tr>
            <td>Destination Port</td>
            <td>16 bits</td>
            <td>Specifies the port number of the receiving process</td>
        </tr>
        <tr>
            <td>Sequence Number</td>
            <td>32 bits</td>
            <td>Used for reliable data transfer</td>
        </tr>
        <tr>
            <td>Acknowledgement Number</td>
            <td>32 bits</td>
            <td>Used for reliable data transfer</td>
        </tr>
        <tr>
            <td>Data Offset</td>
            <td>4 bits</td>
            <td>Specifies the length of the header in 32-bit words</td>
        </tr>
        <tr>
            <td>Reserved</td>
            <td>6 bits</td>
            <td>Reserved for future use</td>
        </tr>
        <tr>
            <td>Control Bits</td>
            <td>6 bits</td>
            <td>Specifies various control information, such as whether the segment contains data, and if it is the
                last segment in a stream</td>
        </tr>
        <tr>
            <td>Window Size</td>
            <td>16 bits</td>
            <td>Specifies the size of the receive window</td>
        </tr>
        <tr>
            <td>Checksum</td>
            <td>16 bits</td>
            <td>A checksum of the entire segment used for error detection</td>
        </tr>
        <tr>
            <td>Urgent Pointer</td>
            <td>16 bits</td>
            <td>Used to specify the location of urgent data, if present</td>
        </tr>
    </tbody>
</table>
<h3>TCP Segment Data</h3>
<p>The data in a TCP segment can vary in size up to the maximum segment size (MSS), which is typically around 1460
    bytes.</p>
<h3>TCP Segment Trailer</h3>
<p>The trailer is optional and typically not used.</p>
<h2>UDP Datagrams</h2>
<p>UDP datagrams consist of a header and data.</p>
<h3>UDP Datagram Header</h3>
<table>
    <thead>
        <tr>
            <th>Field</th>
            <th>Size (bits)</th>
            <th>Description</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Source Port</td>
            <td>16 bits</td>
            <td>Specifies the port number of the sending process</td>
        </tr>
        <tr>
            <td>Destination Port</td>
            <td>16 bits</td>
            <td>Specifies the port number of the receiving process</td>
        </tr>
        <tr>
            <td>Length</td>
            <td>16 bits</td>
            <td>Specifies the total length of the datagram, including the header and data</td>
        </tr>
        <tr>
            <td>Checksum</td>
            <td>16 bits</td>
            <td>A checksum of the entire datagram used for error detection</td>
        </tr>
    </tbody>
</table>
<h3>UDP Datagram Data</h3>
<p>The data in a UDP datagram can vary in size up to 65,535 bytes. UDP is a connectionless protocol and does not
    provide any reliability or flow control mechanisms, so it is often used for applications that require low
    latency and can tolerate packet loss, such as real-time multimedia streaming.</p>
<h2>Conclusion</h2>
<p>In conclusion, network data units are the basic building blocks of network communication. Ethernet frames, Wi-Fi
    frames, IP datagrams, TCP segments, and UDP datagrams each have their own unique format and structure, but they
    all serve the same purpose of transmitting data over a network.</p>
<p>Understanding the structure and format of these network data units is important for troubleshooting network
    issues, optimizing network performance, and designing network protocols and applications.</p>
<p>As networks continue to evolve and new technologies emerge, it's important to stay up to date on the latest
    developments and advancements in network data units. With the increasing importance of data communication in our
    daily lives, knowledge of network data units is essential for anyone working in the field of computer
    networking.</p>









