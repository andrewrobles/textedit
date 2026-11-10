# TextEdit

A small Go CLI (macOS only) that works with TextEdit.

## Usage

```sh
textedit open
textedit open --today
textedit commit
textedit save
```

`open --today` opens (creating it if needed) `~/textedit/yyyy-mm-dd/⠀.rtf` in TextEdit.

`commit` stages all changes in `~/textedit` (which must be a git repo) and commits them with an empty message.

`save` copies `~/textedit/⠀.rtf` to `~/textedit/yyyymmddThhmmss.rtf` (ISO 8601 basic timestamp), resets `⠀.rtf` to the starter content, runs `commit`, then closes and reopens `⠀.rtf` in TextEdit.

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
