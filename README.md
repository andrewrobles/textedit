# TextEdit

A small Go CLI (macOS only) that works with TextEdit.

## Usage

```sh
textedit open
textedit open --today
textedit commit
```

`open --today` opens (creating it if needed) `~/textedit/yyyy-mm-dd/⠀.rtf` in TextEdit.

`commit` stages all changes in `~/textedit` (which must be a git repo) and commits them with an empty message.

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
