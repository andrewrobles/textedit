# TextEdit

A small Go CLI (macOS only) that works with TextEdit.

## Usage

Simply use the `textedit` command to get list of supported commands in ther terminal:

```sh
textedit
```

## Install

Requires Go 1.23 or later.

```sh
go build -o textedit .
sudo mv textedit /usr/local/bin/
```
To try it without installing:

```sh
go run . open
```

## Test

Run all tests from the repository root:

```sh
go test ./...
```

To run a single test with verbose output:

```sh
go test -run TestAppendFile -v
```

The tests read their RTF fixtures from the `test/` directory.
