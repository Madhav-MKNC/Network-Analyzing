<h2>Network Protocol Data Units (NPDUs)</h2>
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