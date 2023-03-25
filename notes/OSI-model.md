# OSI MODEL

<table>
    <thead>
       <tr>
          <th>Layer</th>
          <th>Protocol data unit (PDU)</th>
          <th>Function</th>
          <th>TCP/IP protocols</th>
          <th>Host/Media</th>
          <!-- <th>Host Layers</th>
          <th>Media Layers</th> -->
       </tr>
    </thead>
    <tbody>
       <tr>
          <td align="center"><strong>7</strong><br>Application</td>
          <td rowspan="3" align="center">Data</td>
          <td>High-level protocols such as for resource sharing or remote file access, e.g. HTTP.</td>
          <td>HTTP(s), FTP, SMTP</td>
          <td rowspan="4" align="center"><strong>Host</strong></td>
          <!-- <td></td> -->
       </tr>
       <tr>
          <td align="center"><strong>6</strong><br>Presentation</td>
          <!-- <td>Data</td> -->
          <td>Translation of data between a networking service and an application; including character encoding, data compression and encryption/decryption.</td>
          <td>MIME, SSL/TLS, XDR</td>
          <!-- <td></td>
          <td></td> -->
       </tr>
       <tr>
          <td align="center"><strong>5</strong><br>Session</td>
          <!-- <td>Data</td> -->
          <td>Managing communication sessions, i.e., continuous exchange of information in the form of multiple back-and-forth transmissions between two nodes.</td>
          <td>Sockets</td>
          <!-- <td></td>
          <td></td> -->
       </tr>
       <tr>
          <td align="center"><strong>4</strong><br>Transport</td>
          <td align="center">Segment, Datagram</td>
          <td>Reliable transmission of data segments between points on a network, including segmentation, acknowledgement and multiplexing.</td>
          <td>TCP, UDP, SCTP, DCCP</td>
          <!-- <td></td>
          <td></td> -->
       </tr>
       <tr>
          <td align="center"><strong>3</strong><br>Network</td>
          <td align="center">Packet</td>
          <td>Structuring and managing a multi-node network, including addressing, routing and traffic control.</td>
          <td>IP, IPsec, ICMP, IGMP, OSPF, RIP</td>
          <!-- <td></td> -->
          <td rowspan="3" align="center"><strong>Media</strong></td>
       </tr>
       <tr>
          <td align="center"><strong>2</strong><br>Data link</td>
          <td align="center">Frame</td>
          <td>Transmission of data frames between two nodes connected by a physical layer.</td>
          <td>PPP, SBTV, SLIP</td>
          <!-- <td></td>
          <td></td> -->
       </tr>
       <tr>
          <td align="center"><strong>1</strong><br>Physical</td>
          <td align="center">Bit, Symbol</td>
          <td>Transmission and reception of raw bit streams over a physical medium.</td>
          <td></td>
          <!-- <td></td>
          <td></td> -->
       </tr>
    </tbody>
 </table>
