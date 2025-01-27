# wc-tool

This is a simple command-line utility written in Go that counts lines, words, characters, and bytes in a text file or piped input, similar to the Linux `wc` command.

It is a solution to the first problem in [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc/).


## Steps to Publish

1. Clone this repository.

2. Build the executable:
    ```sh
    go build -o ccwc main.go
    ```

   Alternatively, you can use the `make` command:
    ```sh
    make
    ```
3. Move the executable to a directory in the system's PATH environment:
    ```sh
    sudo mv ccwc /usr/local/bin/
    ```
4. Verify the installation:
    ```sh
    ccwc -h
    ```

## Usage

### Command-line Options

| Option | Description     |
|--------|-----------------|
| -l     | Count lines     |
| -w     | Count words     |
| -c     | Count bytes     |
| -m     | Count characters|

### Examples

#### File Input

To analyze a specific file, pass the file path as an argument:
```sh
ccwc -w myfile.txt
```

#### Piped Input

You can also use the program with piped input:
```sh
cat myfile.txt | ccwc -l
```

#### Default Output

If no flags are provided, the utility outputs the line count, word count, and byte count:
```sh
ccwc myfile.txt
5  25  128 myfile.txt
```