# Explanation

The runner and server (learned from a friend) is a way to startup and shut down a server gracefully.

Nodes communicate via remote procedure calls (RPCs), so the local blockchain network will use RPC. Benefit is that
it abstracts network details, supports interoperability (bridging communication between nodes), serialize & deserialize, etc.

