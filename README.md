# Raft

A Go implementation of the [Raft consensus algorithm](https://raft.github.io/) from scratch.

## Why?

I wanted to understand how distributed consensus works by implementing it. Raft is a more understandable alternative to Paxos, with clear semantics for leader election, log replication, and safety guarantees.

## What's included

- **Leader election** - Nodes elect a leader through voting
- **Log replication** - Safe log synchronization across nodes
- **Persistence** - Term and vote storage, snapshotting support
- **RPC communication** - gRPC-based inter-node communication
- **Configurable logging** - Debug-friendly logging with levels

## Getting started

```bash
make build
make test
```
## Testing

The implementation includes comprehensive tests covering leader election, log replication, and failure scenarios.

```bash
make test
```

## Further reading

I've written blog posts about the implementation process. Check them out on my [blog](https://navgeet.github.io/tags.html#raft-ref).

## Notes

This is a learning project to understand Raft consensus. It implements the core algorithm but isn't production-ready.
