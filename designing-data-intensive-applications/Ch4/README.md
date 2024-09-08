# Chapter 4 super concise summary

## Key Concepts
- Application code and data schemas evolve over time
- Challenge: Schema divergence due to different deployment times
- Need for backward and forward compatibility

## Data Serialization Formats

### Language-specific formats
- Built-in support (e.g., Java Serializable, Python pickle)
- Often problematic due to language lock-in, security issues, poor versioning

### Textual formats (JSON, XML, CSV)
- Human-readable
- Issues with number precision, lack of binary string support
- Inefficient for large datasets

### Binary formats
1. Binary encoding of textual formats (e.g., BSON, MessagePack)
2. Schema-based with code generation (e.g., Protocol Buffers, Thrift)
3. Dynamic schema alongside data (e.g., Avro)

## Modes of Data Flow

### Via databases
- Rolling upgrades require both forward and backward compatibility

### Via service calls
- Web services (REST, SOAP)
- RPC (Remote Procedure Calls)
  - Challenges with network unpredictability, data serialization
  - Newer RPC libraries (gRPC, Thrift, Avro) address some issues

### Via asynchronous messaging
- Uses message brokers (e.g., RabbitMQ, Kafka)
- Benefits: improved reliability, decoupling of sender/receiver

### Actor model
- Programming model for concurrency
- Distributed version uses message brokers for inter-node communication
