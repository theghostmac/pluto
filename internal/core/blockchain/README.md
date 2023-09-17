# Generic elements of a blockchain
The generic elements of a blockchain include the following:
1. The address
2. Transactions
3. Blocks
   A single block is further expounded into different constituent parts.
   All of these make up the execution layer of a blockchain.
   I have already done the implementation of this in this article.
   Let's walk through them in a high level fashion:
4. The Block header - which has the following parts:
   - Previous block header's hash (except the genesis block)
   - Nonce (the contract wallet and the external/user wallet have different nonces)
   - Timestamp (the time of the transaction)
   - Height
   - Merkle root (a bit complex, but we will look at them later).
5. Block body - which has the following part:
   - Transactions.

Headers are important for saving state.

Let's break down what each field in these structs represents:

1. `Header` Struct:
   This struct is used to represent the header of a blockchain block. It contains metadata and essential information about the block.

    - `Version` (uint32):
      This field stores the version number of the block. The version number indicates the format and rules used to interpret the block data. Different versions of the blockchain may have different rules or features, and this field helps identify which version a block adheres to.

    - `PreviousBlockHash` (utils.Hash):
      This field represents the hash of the previous block in the blockchain. Each block in the blockchain contains the hash of the previous block, creating a chain of blocks. This hash is crucial for ensuring the integrity and security of the blockchain, as it connects one block to its predecessor.

    - `Timestamp` (int64):
      The `Timestamp` field stores a Unix timestamp, which is a numeric representation of the time and date when the block was created or mined. It helps maintain the chronological order of blocks within the blockchain.

    - `Height` (uint32):
      The `Height` field indicates the position of the block within the blockchain's sequence. The first block in the blockchain typically has a height of 1, and each subsequent block's height is incremented by one. This field is useful for quickly determining a block's position and for various blockchain-related operations.

    - `Nonce` (uint64):
      The `Nonce` field is often used in the process of mining blocks in a blockchain. Miners change the nonce value repeatedly while attempting to find a hash of the block data that meets certain criteria (e.g., a hash with a specific number of leading zeros). The nonce helps miners perform proof-of-work, which is a critical part of blockchain security.

2. `Block` Struct:
   The `Block` struct represents a complete blockchain block, and it embeds the `Header` struct, meaning it inherits all the fields from the `Header` struct.

    - `Header` (Header):
      This field is an instance of the `Header` struct, which means it includes all the fields (`Version`, `PreviousBlockHash`, `Timestamp`, `Height`, `Nonce`) described above. It allows you to access the block's header information directly through a `Block` instance.

    - `Transactions` ([]transactions.Transactions):
      This field is a slice (an ordered list) that can hold multiple transactions associated with the block. Transactions represent actions or changes to the blockchain, such as transferring cryptocurrency. The `Transactions` field allows you to store and access a list of transactions within the block.

In summary, the `Header` struct stores metadata about a blockchain block, while the `Block` struct encompasses the `Header` and also includes a list of transactions associated with that block. These data structures are fundamental to representing and processing blocks in a blockchain system.


More readable way to create these:
```Go
package blockchain

import (
   "encoding/binary"
   "io"
)

func (h *Header) encodeDecodeBinary(wr io.ReadWriter, binaryFunc func(io.ReadWriter, binary.ByteOrder, interface{}) error) error {
   order := binary.LittleEndian
   err := binaryFunc(wr, order, &h.Version)
   if err != nil {
      return err
   }
   err = binaryFunc(wr, order, &h.PreviousBlockHash)
   if err != nil {
      return err
   }
   err = binaryFunc(wr, order, &h.Timestamp)
   if err != nil {
      return err
   }
   err = binaryFunc(wr, order, &h.Height)
   if err != nil {
      return err
   }
   return binaryFunc(wr, order, &h.Nonce)
}

func (h *Header) EncodeBinary(w io.Writer) error {
   return h.encodeDecodeBinary(w, binary.Write)
}

func (h *Header) DecodeBinary(r io.Reader) error {
   return h.encodeDecodeBinary(r, binary.Read)
}
``