## 💡 What is gRPC?
gRPC stands for Google Remote Procedure Call. It’s a modern, high-performance framework that allows client and server applications to communicate transparently and efficiently.

## ⚙️ Key Concepts:
Term	Description
Remote Procedure Call (RPC)	Call a function on another machine as if it were local.
Protocol Buffers (Protobuf)	A language-neutral binary serialization format used to define the interface and message types.
.proto File	Defines services, RPC methods, and the message format.
Unary RPC	One request → one response (like a normal function call).
Streaming RPC	Streams of data between client and server (uni-directional or bi-directional).
Stub	Auto-generated code (client/server) from the .proto file.

## 🔧 How gRPC Works (Go-specific):
You write a .proto file to define:

The service

The RPC methods (input and output messages)

You compile the .proto file using protoc + Go plugins:

This generates Go code:

*_pb.go: contains message structures

*_grpc.pb.go: contains gRPC server/client interfaces

You implement the server by embedding the generated interface and defining actual logic.

You create a client that calls these methods like regular Go functions.

## 🔁 Workflow Summary
txt
Copy
Edit
[Define .proto] --> [Generate Go Code] --> [Implement Server/Client] --> [Run & Test]

## Server Streaming RPC

In server streaming, the client sends one request, and the server responds with a stream of messages.

📦 Use Case:
Let’s say you want to greet someone multiple times with a delay (like "Hello John #1", "Hello John #2", ...).

## Client Streaming RPC

📦 Use Case:
Imagine the client sends multiple names (e.g., "Alice", "Bob", "Charlie"), and the server replies once:
🗨️ "Hello Alice, Bob, Charlie!"

## Bi-directional Streaming RPC

📦 Use Case:
Imagine a real-time conversation:
The client sends names, and the server immediately replies with a greeting for each — simultaneously, without waiting.

