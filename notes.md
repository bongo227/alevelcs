# The internet

> internet: A global network of network.
###### 
> hypertext transfer protocol (HTTP): The protocol for distributing hypertext documents across the web.


## URL's
> universal resource locator (URL): Address of a resource (e.g. website) on the internet.
#### Structure of a URL
```markdown
http://www.website.co.uk/index.html
  ^     ^     ^     ^  ^    ^
  |     |     |     |  |    |
  |     |     |     |  |   name of the file
  |     |     |     | country code
  |     |     |  top level domain
  |     |   subdomain of co.uk
  |   subdomain of website.co.uk (often the service type)
the protocol
```

> Internet protocol (IP) address: A unique number that identifies a device on a network. Consists of four 4-byte numbers separated with a period.

###### 

## Ports
> Port: Identify's a particular service on the network with a number.
###### 
> POP3: A protocol for receiving emails.
###### 
> SSH: A protocol for remote access to computers.
###### 

#### Well known ports
Services tend to use the same ports as other servers as listed below.
|Port Number|Service|
|---|---|
|21|FTP|
|22|SSH|
|23|Telnet|
|25|SMTP|
|53|DNS|
|80|HTTP|
|110|POP3|
|143|IMAP|

#### Getting a response
When a client requests a service they must send a port for the service to respond back to.

## Network Address Translation (NAT)
> Domain name server (DNS): A server that maps domain names and their IP adresses.

The system that maps private IP's to public IP's is called NAT. This separation of public/private ip addresses means only public IP addresses need to be stored by the DNS and some additional level of protection from the router before packets are are sent to the target device.

#### Loading data from an external server

 - The device creates a packet requesting data from a server. Included is the devices IP adress and port to receive the responce on.
 - The router receives this packet and replaces the internal IP address with it's public address and a generated port number.
 - The router stores the mapping between the private internal IP and port to the external port.
 - Once the response has been fulfilled the router will receive the packet(s) from the server and look at the port number. It then uses its internal map to forward the packet(s) to local device.

> Port forwarding: Exposing a local (non-routable) IP through a router port allowing clients to initiate connections to devices in local networks
###### 
> Socket: Software abstraction that allows defines an endpoint of data flow e.g. Client A opens a socket on port 52707 to receive a stream of packets i.e. a website.

## Subnet's
> Subnet masking: A method of splitting a network into smaller networks, for example the first 4-byte number could represent the building number in a multi-building network
###### 
> Gateway: A device that connects devices to a network with different protocols e.g. An ISP provides a gateway to access the internet, a VPN to connect two company's networks together.

The subnet mask divides IP addresses into the network and host addresses by anding the IP and subnet mask. This allows us to check if a device is on the same network thus we can communicate directly with it.

| Stage | Info |
| --- | --- |
| `120.176.134.32` | IP of workstation |
| `01111000.10110000.01000110.00100000` | Binary equivalent |
| `11111111.11111111.11111111.00000000` | Subnet mask |
| `01111000.10110000.01000110.00000000` | IP & Subnet mask |

## IPv6
The amount of devices connecting to the internet has exploded, consequently we are running out of IPv4 addresses. A new system (IPv6) which uses 128 bits as 8 groups of hex numbers e.g. `13E7:0000:0000:0000:51D5:9CC8:C0A8:6420` allowing many more unique devices.

## Dynamic Host Configuration Protocol (DHCP)
> DHCP: Protocol for allocating unique IP's to local devices.

Static IP addresses are assigned by the router and never change whereas dynamic IP addresses are allocated periodically.

#### DCHP Server
A DHCP server manages a pool of IP adresses.