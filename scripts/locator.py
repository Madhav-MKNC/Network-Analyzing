from scapy.all import *
import requests
import json

ipad = input("Enter Your ip address: ")
#hardcode the ip addr so avoid getting 
sniffer = sniff(filter="udp",count=1)
pkt = sniffer[0]
if str(pkt[IP].src) == str(ipad):
	ipaddr = str(pkt[IP].dst)
if str(pkt[IP].dst) == str(ipad):
	ipaddr = str(pkt[IP].src)


print(ipaddr)
req = requests.get("https://ipinfo.io/"+ipaddr+"/json")
resp = json.loads(req.text)

print ("Location: "+resp["loc"])
print ("region: "+resp["region"])
print ("city: "+resp["city"])
print ("country: "+resp["country"])
print ("postal: "+resp["postal"])
