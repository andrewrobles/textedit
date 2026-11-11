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
