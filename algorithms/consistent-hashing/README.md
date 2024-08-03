# Consistent Hashing Implementation in Zig

This project implements a consistent hashing algorithm in Zig, demonstrating how to distribute keys across a set of servers in a way that minimizes redistribution when servers are added or removed.

## Features

- Add and remove servers dynamically
- Distribute keys across available servers
- Visualize the hash ring in the command line
- Efficient key lookup
- Minimal key redistribution when changing the server set

## Requirements

- Zig compiler (tested with version 0.10.0)

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/consistent-hashing-zig.git
   cd consistent-hashing-zig
   ```

2. Build & Run the project:
   ```
   zig run consistent-hashing.zig
   ```


The program provides an interactive command-line interface with the following commands:

- `add <server_name>`: Add a new server to the hash ring
- `remove <server_name>`: Remove a server from the hash ring
- `get <key>`: Find which server is responsible for a given key
- `visualize`: Display an ASCII representation of the current hash ring
- `quit`: Exit the program

## Visualization by ByteByteGo

<img width="975" alt="Screenshot 1446-01-28 at 1 18 04 PM" src="https://github.com/user-attachments/assets/28068f35-a0e3-4a0c-ad7c-7db463f28a1f">
