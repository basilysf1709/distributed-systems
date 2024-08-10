# Structured Go Key-Value Store

This is a simple key-value store implemented in Go with a command-line interface (CLI) and JSON file persistence. The project is structured to allow easy expansion and addition of new features.

## Project Structure

```
kvstore/
├── cmd/
│   └── kvstore/
│       └── main.go           # Entry point of the application
├── internal/
│   ├── store/
│   │   ├── store.go          # Internal implementation of the store
│   │   └── json_persistence.go # JSON file persistence implementation
│   └── cli/
│       └── cli.go            # Command-line interface implementation
├── pkg/
│   └── kvstore/
│       └── kvstore.go        # Public API for the key-value store
├── go.mod
└── README.md
```

## Features

- Set, get, and delete key-value pairs
- Persistent storage using a JSON file
- Command-line interface for easy interaction
- Thread-safe operations
- Modular design for easy expansion

## Installation

1. Ensure you have Go installed on your system. If not, download and install it from [golang.org](https://golang.org/).

2. Clone this repository:
   ```
   git clone https://github.com/yourusername/kvstore.git
   cd kvstore
   ```

3. Build the program:
   ```
   go build -o kvstore ./cmd/kvstore
   ```

## Usage

The program supports three main operations: set, get, and delete.

### Setting a Key-Value Pair

```
./kvstore set <key> <value>
```

Example:
```
./kvstore set name "John Doe"
```

### Getting a Value

```
./kvstore get <key>
```

Example:
```
./kvstore get name
```

### Deleting a Key-Value Pair

```
./kvstore delete <key>
```

Example:
```
./kvstore delete name
```

## Data Persistence

The key-value pairs are stored in a file named `kvstore.json` in the same directory as the program. This file is automatically created if it doesn't exist and is updated after each set or delete operation.

## Error Handling

The program includes error handling:

- If a file operation fails, an error message will be displayed.
- If you try to get a non-existent key, the program will inform you that the key was not found.
- If you use the program incorrectly, it will display usage instructions.

## Extending the Project

To add new features:

1. Implement new persistence methods in `internal/store/`
2. Add new commands in `internal/cli/cli.go`
3. Extend the public API in `pkg/kvstore/kvstore.go` if necessary