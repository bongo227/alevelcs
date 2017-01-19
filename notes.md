# AQA Computing
By Ben Sheffield

---
# Queues and stacks

## Stacks

> LIFO: Data structure where the last item added is the first one out.
###### 
> Stack: LIFO structure

#### Pop of the stack
Let `stackPointer` be the index of the top of the stack and `stackArray` the array used to store the data items.
```
FUNCTION pop()
	IF stackPointer > 0 THEN
		data <- stackArray[stackPointer]
		stackPointer <- stackPointer - 1
		RETURN data
	ELSE
		ERROR "Stack is empty"
	END IF
END FUNCTION
```

#### Push on to the stack
Let `stackSize` be the maximum size of the stack.
```
FUNCTION push(item)
	IF stackPointer < stackSize THEN
		stackPointer <- stackPointer + 1
		stackArray[stackPointer] <- item
	ELSE
		ERROR "Stack is full"
	END IF
END FUNCTION
```

#### Stack Frames
Stacks are used to to store data for use in the subroutines of running programs.

> Stack frame: Data pertaining to a subroutine call
##### 
> Call stack: A stack of all active stack frames

#### Structure of a stack frame
| Stack frame | Description |
| --- | --- |
| Parameters | Values of the parameter passed to the procedure |
| Return address | Address of where to jump after procedure returns |
| Local variables | Any local variables allocated in the procedure |

## Queues
> FIFO: A data structure where the first item in is also the first item out.
###### 
> Queue: A FIFO data structure.
###### 
> Linear queue: A queue implemented as a list
###### 
> Circular queue: A queue implemented as a ring where the front and rear pointers can wrap around to the start of the array
###### 
> Priority queue: A queue where the first item to leave is the one with the highest priority
 
#### En-queue circular queue
Let `rearPointer` be the index of the rear of the queue and `queueSize` be the max size of the queue.
```
FUNCTION enqueue(item)
	rearPointer <- (rearPointer + 1) % queueSize
	queueArray[rearPointer] <- item
END FUNCTION
```

#### De-queue circular queue
Let `frontPointer` be the index of the front of the queue.
```
FUNCTION dequeue()
	data <- queueArray[frontPointer]
	frontPointer <- (frontPointer + 1) % queueSize
	RETURN data
END FUNCTION
```
---
# Graphs and trees
> Graph: A model of the relationship between objects.
###### 
> Edge: Relationship between to nodes.
###### 
> Node: Object in a graph.
###### 
> Weighted graph: Graph where the edges have a particular weight.
###### 
> Undirected graph: Relation ship between nodes can go in either direction.
###### 
> Directed graph: One way relationship between nodes.


## Adjacency list
> Adjacency list: A data structure that stores a list of nodes and their adjacent node

```markdown
A ---- B ---- E
 \   /  \   /
  \ /    \ / 
   C ---- D
```
To turn this graph into an adjacency list just write down each node then a list of nodes that connect to them as below.

| Node | Adjacent nodes |
| --- | --- |
| A | B, C |
| B | A, C, E |
| C | A, B, D |
| D | C, E |
| E | B, D |

In a directed graph only the connections from each node to another are listed as adjacent.
In a weighted graph the value of each edge is also listed like: B, 20, C, 10.

## Adjacency matrix
> Adjacency list: A 2d array where 1's represent an edge and 0's represent no edge.

Using the same graph as above the following adjacency matrix is produced
| | A | B | C | D | E |
| --- | --- | --- | --- | --- | --- |
| **A** | 0 | 1 | 1 | 0 | 0 |
| **B** | 1 | 0 | 1 | 0 | 1 |
| **C** | 1 | 1 | 0 | 1 | 0 |
| **D** | 0 | 0 | 1 | 0 | 1 |
| **E** | 0 | 1 | 0 | 1 | 0 |

In a directed graph only the connections from each node to another are 1's.
In a weighted graph the value of the edge is listed instead of a 1, where their is no edge the value is ∞.

## Adjacency list vs matrix
| | List | Matrix |
| --- | --- | --- |
| Storage | Less memory since it only stores data when an edge exsits | More memory because it stores a value for every possible combination of edges |
| Processing | Long processing because the list must be passed to find if a edge exsists | Short processing because matrix directly stores wether an edge exsists or not |
| When is it suitable | Not many edges (sparse graph) | Many edges (dense graph) | 

## Trees
> Tree: A graph with no loops
###### 
> Root: The node all other nodes branch from.
###### 
> Parent: A node with children nodes
###### 
> Child: A node attached to a parent node
###### 
>  Leaf: A node with no children

#### Binary search tree
> Binary Tree: A tree where parents only have up to two children

On type of tree is a binary search tree which makes finding data quicker. To append to a BST:
1. Start at the root node, if the value of the node is less than the root node move left else move right
2. Continue for until you reach a child node
3. If the value of node is less than the child node append it left of the child node, else append it to the right.

---
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

