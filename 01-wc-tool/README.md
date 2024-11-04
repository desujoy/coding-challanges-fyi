# ccwc - A Custom Implementation of the `wc` Command

## Overview

`ccwc` is a simplified implementation of the Unix `wc` (word count) command line tool. This tool allows you to count the number of bytes, lines, words, and characters in a given file or from standard input, following the Unix Philosophy of building simple, modular tools.

## Features

- **Byte Count (`-c`)**: Outputs the number of bytes in the input file.
- **Line Count (`-l`)**: Outputs the number of lines in the input file.
- **Word Count (`-w`)**: Outputs the number of words in the input file.
- **Character Count (`-m`)**: Outputs the number of characters in the input file (supports multibyte characters if your locale allows).
- **Default Mode**: When no options are provided, `ccwc` displays the byte, line, and word counts.
- **Standard Input Support**: If no filename is provided, `ccwc` reads from standard input.

## Installation

To install `ccwc`, ensure you have [Go](https://golang.org/dl/) installed on your system. Then, clone this repository and build the tool:

```sh
go build -o ccwc
```

This will generate an executable named `ccwc` in the current directory.

## Usage

The basic usage of `ccwc` is:

```sh
./ccwc [OPTIONS] [FILE]
```

### Options

- `-c`: Print the byte count.
- `-l`: Print the line count.
- `-w`: Print the word count.
- `-m`: Print the character count.

### Examples

1. **Count bytes in a file**:
   ```sh
   ./ccwc -c test.txt
   ```

   **Output**:
   ```
   342190 test.txt
   ```

2. **Count lines in a file**:
   ```sh
   ./ccwc -l test.txt
   ```

   **Output**:
   ```
   7145 test.txt
   ```

3. **Count words in a file**:
   ```sh
   ./ccwc -w test.txt
   ```

   **Output**:
   ```
   58164 test.txt
   ```

4. **Count characters in a file**:
   ```sh
   ./ccwc -m test.txt
   ```

   **Output**:
   ```
   339292 test.txt
   ```

5. **Run without options** (defaults to `-c`, `-l`, `-w`):
   ```sh
   ./ccwc test.txt
   ```

   **Output**:
   ```
   7145 58164 342190 test.txt
   ```

6. **Read from standard input**:
   ```sh
   cat test.txt | ./ccwc -l
   ```

   **Output**:
   ```
   7145
   ```

## Code Overview

The `ccwc` tool is implemented in Go and uses `bufio` to efficiently read files and standard input. The main function handles:

- Parsing command line arguments using `flag`.
- Opening the specified file or reading from standard input.
- Counting bytes, lines, words, and characters using a custom function `getFileCounts`.

### Key Functions

- **`getFileCounts(reader *bufio.Reader)`**: Reads the input and returns the counts for bytes, lines, words, and characters.
- **`main()`**: The entry point that sets up command-line parsing, handles file reading, and prints results based on specified options.