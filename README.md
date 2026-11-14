# TextEdit

A small Go CLI (macOS only) that works with TextEdit.

## Usage

Simply use the `textedit` command to get list of supported commands in ther terminal:

```sh
textedit
```

## Install

Requires Go 1.23 or later.

### From GitHub

Install straight from the repository:

```sh
go install github.com/andrewrobles/textedit@latest
```

This puts the `textedit` binary in `$(go env GOPATH)/bin` (usually `~/go/bin`), so make sure that directory is on your `PATH`:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

### From source

Clone the repository, then build and install it:

```sh
git clone https://github.com/andrewrobles/textedit.git
cd textedit
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
