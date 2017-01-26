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
# Tree traversal

> **Depth first**: Goes down the left most branch before backtracking.
> **Breadth first**: Explores the closest node before exploring deeper nodes.
> **Pre-order**: Visits the root then, transverses the left sub-tree then the right sub-tree.
> **In-order**: Transverses the left sub-tree, visits the root then, then transverses the right sub-tree.
> **Post-order**: Transverses the left sub-tree then the right sub-tree and then visits the root node. 

#### Traversing example
```markdown
        a
       / \
      /   \
     b     e
    / \
   /   \
  c     d
```

Pre-order: a, b, c, d, e
In-order: c, b, d, a, e
Post-order: c, d, b, e, a

#### Quick method
``` markdown
  pre-order > N < post-order
              ^
          in-order
```
Draw a line around the binary tree, you will visit each node three times from the left, bottom and right. Depending on how you transverse the graph you recorded each node when you visit the corresponding side.

---
# Dijkstra's algorithm

```markdown
    5      7
A ----- B ----- E
 \  9 /   \    / 10
2 \  /   1 \  /  
   C ------ D
       3
```

1. Start at `A`
2. Now compute the total weight of the current path (`A`) plus the weight of its neighbors:
	`A-B`: 5
	`A-C`: 2
3. Follow the shortest root and repeat
	`A-C-B`: 2 + 9 = 11
	`A-C-D`: 2 + 3 = 5
4. If you have a root with the same start and end point i.e. `A-B` and `A-C-B`, discard the longest.
5. Repeat until all vertices have been visited

### Table notation

| Step | Vertex | A | B | C | D | E |
| --- |
| 1 | `A` | $\color{red}{0_A}$ | $5_A$ | $2_A$ | $\infty_A$ | $\infty_A$ |
| 2 | `C` | $\color{red}{0_A}$ | $5_A$ | $\color{red}{2_A}$ | $5_C$ | $\infty_C$ |
| 3 | `D` | $\color{red}{0_A}$ | $5_A$ | $\color{red}{2_A}$ | $\color{red}{5_C}$ | $15_D$ |
| 4 | `B` | $\color{red}{0_A}$ | $\color{red}{5_A}$ | $\color{red}{2_A}$ | $\color{red}{5_C}$ | $12_B$ |
| 5 | Done | $\color{red}{0_A}$ | $\color{red}{5_A}$ | $\color{red}{2_A}$ | $\color{red}{5_C}$ | $\color{red}{12_B}$ | 


- Record the distance between `A` and any nodes its connected to. 
- The subscript denoted which node that value was recorded from so you can trace back the shortest route. 
- Color in red the current node, all nodes in red are known to be the shortest path to that point.
- Move to closest non-visited node and repeat.
- When their is no more nodes to visit record the final values in a separate row.

---
# Reverse Polish Notation
> **Infix**: Operators are within their operands i.e. `5 + 3`
> **Reverse Polish Notation**:   Operators come before their operands i.e. `+ 5 3`


### Converting in-fix to pre-fix

- Number => Push to out queue
- Operator => If the operator stack has a lower precedence (BIDMAS) push in to the stack, else pop the stack onto the out queue until the precedence is less than or equal to the precedence of the current token
- Left parenthesis => Push to the stack
- Right parenthesis => Pop the stack onto the out queue until you reach the left parenthesis, then pop the left parenthesis.
- End => Pop the rest of the stack

#### Example

Given the expression `2 * (4 + 3) + 1`
| Symbol | Out queue | Stack | Note |
| --- | --- | --- | --- |
| `2` | `2` |  | Queued `2`  |
| `*` | `2` | `*` | Pushed `*` onto the stack |
| `(` | `2` | `*`, `(` | Pushed `(` onto the stack | 
| `4` | `2 4` | `*`, `(` | Queued `4` |
| `+` | `2 4` | `*`, `(`, `+` | Pushed `+` onto the stack |
| `3` | `2 4 3` | `*`, `(`, `+` | Queued `3` |
| `)` | `2 4 3 +` | `*` | Popped and queued up to and including the `(` |
| `+` | `2 4 3 + *` | `+` | Popped and queued `*`, pushed `+` onto the stack |
| `1` | `2 4 3 + * 1` | `+` | Queued `1` |
| | `2 4 3 + * 1 +` | | Popped and queued `+` |


---
# Finite state machines

### State transition diagrams
State transition diagrams use circles the represent state and arrows to represent how input changes the current state to another. The **accepting state** is the final state where the goal is met.

#### Example
Consider a machine that accepts 1p and 2p coins and requires 2p in total.
```markdown
-> (S0) --1--> (S1) --1--> ((S2))
     |                       ^
     |                       |
     -----------2-------------
```

### Mealy machine
A subsection of FSMs that produce some output. Each arrow is labled with some input and output. `1/0` means accept `1` and output `0`. 

#### Example
This mealy machine performs a right shift
[![148-3.jpg](https://s23.postimg.org/k4kryfqln/148_3.jpg)](https://postimg.org/image/mlwj5pahz/)

### State transistion table
Another way to represent a FSM would be in table form.

#### Example
Consider the right shift diagram.
| Input | Current state | Output | Next state |
| --- | --- | --- | --- |
| 1 | S0 | 0 | S1 |
| 0 | S0 | 0 | S2 |
| 1 | S1 | 1 | S1 |
| 0 | S1 | 0 | S2 |
| 1 | S2 | 0 | S2 |
| 0 | S2 | 1 | S1 |


# Turing machine
A Turing machine consists of some **tape** and a **read/write head**. The tape is split into an infinite amount of **cells** which the read/write head can manipulate. Once it has completed the machine enters the **halting state**. The **alphabet** is a list of symbols which can be stored in each cell and a **transition function** describes what should be stored in a cell and weather to move left or right.

### Transition rules
One way of describing a Turing machine is a sequence of transition rules such as: `δ(S0, 0) = (S1, 1, R)`. The format is a `δ` followed by the current state and input symbol. On the right is the next state, value to write and the direction to move. If the head is in that state and it reads the input symbol the rule is applied.

### Universal machine
A Turing machine can solve any computable problem however every process requires its own Turing machine. A universal machine is a Turing machine which simulates a Turing machine given a description of that Turing machine and its inputs.

```markdown
                             +----------+
 Description of machine ---> | Univeral | ---> Output
                  Input ---> | machine  |
                             +----------+
```

---
# Regular and context-free languages

### Regular expression
Regular expressions are a simple language which searched a set of characters.
| Regular expression | Meaning | Matches |
| --- | --- | --- |
| `a\|b\|c` | a or b or c | "a", "b", "c" |
| `abc` | a and b and c | "abc" |
| `a*bc` | zero or more a's followed by b and c | "bc", "abc", "aaabc" |
| `(a\|b)c` | a or b and c | "ac", "bc" |
| `a+bc` | one or more a and b and c | "abc", "aaabc" |
| `a?bc` | zero or one a and b and c | "bc", "abc" |

### Standard expressions
The following are some common extensions to regular expressions.
| Regular expression | Meaning | Matches |
| --- | --- | --- |
| `.ole` | wild card and o and l and e | "hole", "mole" |
| `[mh]` | m or h | "m", "h" |
| `[\^m]` | not m | "a", "h", "b" |
| `{m, 4}` | m and m and m and m | "mmmm" | 

### Context free language
Describes the syntax of a language when the syntax is more complex. One way of describing a context-free language would be Backus-Baur form:

```
<digit> ::= 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9
<integer> ::= <digit> | <digit> <integer>
```

Each element is described by a combination of other elements. `<digit>` is a number 0-9 and the `integer` is either a single digit or a digit and an integer which would match numbers of any length.

# The binary number system

### Adding binary numbers
Since its base-2 just add the 1's, if you get >1, carry the 1.
```markdown
 0 0 1 1 0 0 1 0 +
 1 0 1 1 0 1 0 1
 ---------------
 1 1 1 0 0 1 1 1 
 ---------------
   1 1
```

### Multiply binary numbers
For each 1 on the right of the `x` write out the left hand side from right to left starting at the position of the 1 on the right.
```markdown
    1 1 0 1 1 x
          1 1
-------------
    1 1 0 1 1
  1 1 0 1 1 0
-------------
1 0 1 0 0 0 1
-------------
  1 1 1 1
```

### Make a binary number negative
To convert `01100110`'s sign just move right to left until you reach a `1`, move past the `1` then flip the rest of the bits so
```
0 1 1 0 0 1 1 0

1 0 0 1 1 0 1 0
```

### Floating point
Floating point numbers have a mantissa and an exponent just like a denaty number in standard form. For example:
| $-1$ | $\frac{1}{2}$ | $\frac{1}{4}$ | $\frac{1}{8}$ | $\frac{1}{16}$ | $\frac{1}{32}$ | $\frac{1}{64}$ | $\frac{1}{128}$ | \| | $-8$ | $4$ | $2$ | $1$ |
| --- |
| 0 | 0 | 0 | 0 | 1 | 1 | 0 | 0 | \| | 0 | 0 | 1 | 1

is the same as $\frac{3}{32} \times 2^3$ or $0.75$

#### Normalization
To convert a number to floating point form we must normalize it. We must make the shift it left or right until the first bit is the sign bit followed by the floating point.

| Number | Comment |
| --- | --- |
| `-15.75` | Denary number |
| `10001.11` | Fixed point form |
| `1.000111 1100` | Floating point form |

### Errors

#### Absolute error
$$ \text{expected value} - \text{actual value} $$
For example: $6.95 - 6.9365 = 0.0125$

#### Relativer error
$$ \frac{\text{absolute error}}{\text{expected value}} $$
For example: $\frac{0.0125}{6.95} = 0.001798561151$

---
# The internet

> internet: A global network of networks.
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

