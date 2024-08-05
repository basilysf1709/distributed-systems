## Hinted Handoff Real life scenario

The program simulates a small distributed system with three nodes. It demonstrates:

1. Writing data to a node
2. Replicating data across nodes
3. Handling a node going offline
4. Using hinted handoff to store data intended for an offline node
5. Processing hints when a node comes back online

## Features

- Simulates a basic distributed system with three nodes
- Implements basic write and read operations
- Demonstrates data replication across nodes
- Implements a simple hinted handoff mechanism
- Provides an interactive, step-by-step demonstration of the algorithm

## Requirements

- Zig compiler (tested with version 0.11.0)

## Building and Running

To build and run the project:

1. Ensure you have Zig installed on your system.
2. Navigate to the project directory.
3. Run the following commands:

```
zig run hinted-handoff.zig
```

## Code Structure

- `Node`: Struct representing a node in the system
- `DataItem`: Struct representing a piece of data
- `writeData`: Function to simulate writing data to a node
- `replicateData`: Function to simulate data replication across nodes
- `hintedHandoff`: Function to implement the hinted handoff mechanism
- `processHints`: Function to process hints when a node comes back online

## Limitations

This is a simplified implementation for educational purposes. In a real-world scenario:

- The system would handle more complex failure scenarios
- There would be more sophisticated mechanisms for choosing hint storage locations
- The system would include features like consistency levels and read repair

## Further Development

Potential areas for expansion:

- Implement more complex node selection strategies for hint storage
- Add support for varying consistency levels
- Implement a gossip protocol for node status updates
- Add support for read operations and read repair
