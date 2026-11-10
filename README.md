# TextEdit

A small Go CLI (macOS only) that works with TextEdit.

## Usage

```sh
textedit open
textedit open --today
```

`open --today` opens (creating it if needed) `~/textedit/yyyy-mm-dd/⠀.rtf` in TextEdit.

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
