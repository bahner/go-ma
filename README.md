Go [間] reference implementation
===

Provided are the tools and and structs to interact with [SPACE] through [間]. However this library can be used for any purpose. It is not tied to [SPACE] by design.

Development specs will reside in <https://github.com/bahner/ma>. But the specs are mostly written as an afterthought. This code is the spec.

To understand the need for 間 I commend [Alan Kay's talk at OOPSLA 1997][kayma] to you. It is brilliant, thought provoking, inspiring and highly entertaining.

A success criterium for this project would be if he looked at the code and described it as: *Mostly harmless*.

So what is it? It's almost, but not quite, entirely unlike [DID][did] and [DIDComm][didcomm] over [libp2 pubsub][pubsub]. Another important point made by Kay is that each [object should have a URL and an IP][kayurl]. 間 provides that by leveraging DID's which express IPNS identities and uses pubsub over [IPFS][ipfs] ..

The envelope contains the encrypted message and the key. The message is a DID Document. The key is the key to unlock the message. The DID Document is signed by the sender. And is structured, so the receiver can parse it for it's intended purpose.

Go language
---

I intended to create the prototype entirely in Elixir, but [libp2p] in Elixir is not as mature as in Go.

I have never written a single line of Go before, so this is a learning experience for me. I wrote the required [IPFS][exipfs], [IPNS][exipns] and [IPLD][exipld] libraries for Elixir. But a libp2p implementation is not feasible in the foreseeable future. At least for one person. So I decided to do it in Go.

These structs and functions are the result of that decision. They will be consumed by another project - [go-space] - which is an erlang node that will be used to communicate with the Go 間 tools and integrate with the MOO that I am building in Elixir.

did
---

The did sub-package contains the structs and functions to build and work with DID documents. I did look at other did-implementations, but they were too complicated for me, so I decided to write my own.

These packages do not provide for publication or retrieval of the documents. Live interactions with these entities should be consumed by other software using this library. It is intended to be as lightweight, clean and stable as possible.

I will not add unneeded functionality to this package at the moment. It started as more of a full
implementation of the DID spec, but I have decided to keep it as simple as possible.

If need be it should be relatively easy to extend, as it's been endowed with a version number.

Message
---

The message struct is also meant to be as stable as possible. It is the struct
that is used to create messages to be sent on 間. The envelope consists of a an encrypted payload and the key to unlock it.

As noted above the code is the spec for now, but we'll get to that. RSN.

2023-11-25: bahner

[did]: <https://www.w3.org/TR/did-core/> "Decentralized Identifiers (DIDs) v1.0"
[didcomm]: <https://identity.foundation/didcomm-messaging/spec/> "DIDComm Messaging v1.0"
[exipfs]: <https://hex.pm/packages/ex_ipfs> "Elixir IPFS"
[exipld]: <https://hex.pm/packages/ex_ipfs_ipld> "Elixir IPLD"
[exipns]: <https://hex.pm/packages/ex_ipfs_ipns> "Elixir IPNS"
[go-space]: <https://github-com/bahner/go-space> "Go SPACE"
[ipfs]: <https://ipfs.tech> "InterPlanetary File System"
[kayma]: <https://www.youtube.com/watch?v=oKg1hTOQXoY&t=2268s> "The space between objects."
[kayurl]: <https://www.youtube.com/watch?v=oKg1hTOQXoY&t=2582s> "Every object should have a URL and and IP"
[libp2p]: <https://libp2p.io> "libp2p"
[pubsub]: <https://docs.libp2p.io/concepts/pubsub/overview/> "libp2p pubsub"
[SPACE]: <https://github.com/bahner/space> "SPACE"
[間]: <https://github.com/bahner/ma> "ma specs"
