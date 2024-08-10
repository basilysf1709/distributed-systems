# Go Key-Value Store with CLI

This is a simple key-value store implemented in Go with a command-line interface (CLI) and JSON file persistence.

## Features

- Set, get, and delete key-value pairs
- Persistent storage using a JSON file
- Command-line interface for easy interaction
- Thread-safe operations

## Installation

1. Ensure you have Go installed on your system. If not, download and install it from [golang.org](https://golang.org/).

2. Clone this repository:
   ```
   git clone https://github.com/basilysf1709/go-kv-store.git
   cd kvstore
   ```

3. Build the program:
   ```
   go build -o kvstore
   ```

## Usage

The program supports three main operations: set, get, and delete.

### Setting a Key-Value Pair

To set a key-value pair:

```
./kvstore set <key> <value>
```

Example:
```
./kvstore set name "John Doe"
```

### Getting a Value

To retrieve a value for a given key:

```
./kvstore get <key>
```

Example:
```
./kvstore get name
```

### Deleting a Key-Value Pair

To delete a key-value pair:

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

The program includes basic error handling:

- If a file operation fails, an error message will be displayed.
- If you try to get a non-existent key, the program will inform you that the key was not found.
- If you use the program incorrectly, it will display usage instructions.

## Limitations

- All values are stored as strings.
- The entire store is read from and written to the file for each operation, which may be inefficient for large datasets.
- There's no built-in way to list all keys or values.

## Future Improvements

- Add a 'list' command to display all key-value pairs.
- Implement backup and restore functionality.
- Add support for different data types.
- Create an interactive mode (REPL) for continuous operation.
- Allow specifying a custom file path for the JSON storage.