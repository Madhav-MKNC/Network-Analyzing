using Sockets
using Base.Iterators
using Bitwise

function discover_hosts()
    # Get the IP address and netmask of the network interface
    iface = "en0"  # Change to the name of your network interface
    ip = strip(chomp(read(`ifconfig $iface | grep 'inet ' | awk '{print $2}'`, String)))
    netmask = strip(chomp(read(`ifconfig $iface | grep 'inet ' | awk '{print $4}'`, String)))

    # Calculate the network address and broadcast address
    ip_octets = split(ip, ".")
    netmask_octets = split(netmask, ".")
    network_octets = join.([string(bitwise_and(parse(Int, ip_octets[i]), parse(Int, netmask_octets[i])) & 0xFF) for i in 1:4], ".")
    network = network_octets
    broadcast_octets = join.([string(bitwise_or(parse(Int, network_octets[i]), bitwise_not(parse(Int, netmask_octets[i])) & 0xFF) & 0xFF) for i in 1:4], ".")
    broadcast = broadcast_octets

    # Send ARP request packets to all hosts on the network
    hosts = []
    for i in 1:254
        target_ip = "$network.$i"
        s = UDPSocket()
        setsockopt(s.sockfd, IPPROTO_IP, IP_HDRINCL, true)
        bind(s, iface)
        arp_request = zeros(UInt8, 42)
        arp_request[1:6] .= 0xFF
        arp_request[7:12] .= reinterpret(UInt8, parse.(UInt8, split("00:00:00:00:00:00", ":")))
        arp_request[13:14] .= 0x0806
        arp_request[15:16] .= 0x0001
        arp_request[17:18] .= 0x0800
        arp_request[19:20] .= 0x06
        arp_request[21:22] .= 0x0001
        arp_request[23:24] .= 0x0000
        arp_request[25:30] .= reinterpret(UInt8, parse.(UInt8, split("00:00:00:00:00:00", ":")))
        arp_request[31:34] .= reinterpret(UInt8, parse.(UInt8, split(ip, ".")))
        arp_request[35:40] .= reinterpret(UInt8, parse.(UInt8, split("00:00:00:00:00:00", ":")))
        arp_request[41:44] .= reinterpret(UInt8, parse.(UInt8, split(target_ip, ".")))
        sendto(s, arp_request, (target_ip, 0x0806))
        try
            packet = recv(s)
            if packet[13:14] == reinterpret(UInt16, 0x0806) && packet[21:22] == reinterpret(UInt16, 0x0002)
                mac = join([string(hex(x, 2)) for x in packet[23:28]], ":")
                push!(hosts, (target_ip, mac))
            end
        catch e
        end
        close(s)
    end

    return hosts
end

if abspath(PROGRAM_FILE) == @__FILE__
    hosts = discover_hosts()
    for host in hosts
        println("IP address: $(host[1]), MAC address: $(host[2])")
    end
end
