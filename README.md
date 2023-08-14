# Pluto is a blockchain implementation in Go.
What is a blockchain? [Here](https://en.wikipedia.org/wiki/Blockchain).

Pluto is a blockchain I build from scratch with Go. I am re-exploring blockchain technology for the second time,
but this time, I am doing it properly.

# Roadmap
- [x] Network Layer: RPC communication
  - [x] Create the local transport network model
  - [x] Implement methods for the local transport of payload from one node to another
  - [x] Create an RPC server for communication between two nodes
  - [x] Write tests for the implemented transport methods

![Tests passes](testPasses.png)

- [ ] Execution Layer: Blocks and Transactions
  - [x] Block model
  - [x] Encode and Decode data from the head and the block
  - [x] Write tests for blocks
  - [ ] Transactions model