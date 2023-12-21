# Blockchain Engineering from first principles

This article is for people who already know what the blockchain is, and probably already build software
around it. It is inspired by my journey to relearn blockchain technology beyond the surface. While I do
this, I try to build a blockchain from scratch using the Go programming language.

## Introducing blockchain

Blockchain is a peer-to-peer, distributed ledger that is cryptographically secure, append-only, immutable,
and update-able only via consensus among peers.\

## Definition of terms

1. Peer-to-peer means there is no central controller of the network, and all participants (called nodes)
talk to each other directly.
2. Distributed ledger means that the ledger is spread across the network among all peers in the network,
each peer holding a copy of the complete ledger.
3. Cryptographically secure means that the ledger is secured from tampering and misuse using
cryptographic algorithms.
4. Append-only means that data can only be added, and not modified or tampered. Also, the data
is added to the blockchain in time-sequential order.
5. Update-able via consensus means that a consensus mechanism is used to enable decentralization
on the blockchain.

From the following fundamental definition of blockchain, it is obvious that a blockchain
engineer must be versed in:

1. distributed systems - I recommend Distributed Systems for Practitioners book.
2. writing immutable code.
3. cryptography - I recommend Cryptographic Algorithms
4. mathematics, and
5. computer networking.

In addition to all of these recommendations, you can add Wikipedia, and any other resources you find.
It's helpful to listen to experts talk, so videos might be good too.

## Blockchain Architecture

Blockchain is a network with different layers. It is similar to HTTP, FTP, etc., which
runs on the TCP/IP model. Just as the TCP/IP networking model has 4 layers, the blockchain networking model
has 6 layers:

- Application:
  - Smart contracts
  - Decentralized applications
  - Decentralized autonomous organizations
  - Autonomous agents
- Execution:
  - Virtual machines
  - Blocks
  - Transactions
- Consensus:
  - State machine replication
  - Proof-based consensus
  - Traditional Byzantine fault-tolerant protocols
- Cryptography:
  - Public key cryptography
  - Digital signatures
  - Hash functions
- P2P:
  - Gossip protocols / Epidemic protocols
  - Routing protocols
  - Flooding protocols
- Network:
  - The internet
  - TCP/IP

From the list, it is obvious that the Network layer is the lowest layer. So in building a
blockchain from scratch, it is best to start from the Network layer.

### The Network layer

Blockchain nodes communicate using remote procedure calls.

Network connections are also built on top of hardware that will also fail at some
point and we should design our systems accordingly.

#### Basics and Mechanics of the Network layer

The mechanics of the blockchain can be broken down into the following:

1. **Node Communication**: In a blockchain, nodes communicate via a peer-to-peer network protocol. While it's not exactly like a normal web server, you can think of it as a distributed network where nodes share information directly.

2. **Initial Information**: New nodes can bootstrap by connecting to existing nodes, which provide them with information about the blockchain's current state.

3. **Coin Transactions**: People buy coins by interacting with the blockchain's protocol. In some cases, they might exchange value through these interactions.

4. **Storing Data**: The blockchain's data structure stores transaction data in blocks, which are linked together in a chain. Each node maintains a copy of this data.

5. **RPC Endpoints**: RPC endpoints allow external applications to communicate with nodes for retrieving data or submitting transactions.

6. **Central Servers**: While blockchains are decentralized, the concept of a central server doesn't apply in the same way. However, some blockchains might have centralized components like explorer websites that display blockchain data.

7. **Data Propagation**: Nodes in a blockchain network communicate to propagate new transactions and blocks. This is done through a consensus protocol, ensuring that all nodes eventually agree on the state of the blockchain.

### The Execution layer

I particularly started the execution layer in order to have considerably real-world payload to feed the nodes for their communication. To
understand the execution layer, you must first understand the generic elements of a blockchain:

#### Generic elements of a blockchain

The generic elements of a blockchain include the following:

1. The address
2. Transactions
3. Blocks
   A single block is further expounded into different constituent parts.
   All of these make up the execution layer of a blockchain.
   I have already done the implementation of this in this article.

Let's walk through them in a high level fashion:

1. The Block header - which has the following parts:

   - Previous block header's hash (except the genesis block)
   - Nonce (the contract wallet and the external/user wallet have different nonces)
   - Timestamp (the time of the transaction)
   - Height
   - Merkle root (a bit complex, but we will look at them later).

2. Block body - which has the following part: Transactions.

Headers are important for saving state.
