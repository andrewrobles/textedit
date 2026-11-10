package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// The file name is a single braille blank character (U+2800) followed by ".rtf".
const fileName = "\u2800.rtf"

func main() {
	today := false
	switch {
	case len(os.Args) == 2 && os.Args[1] == "open":
	case len(os.Args) == 3 && os.Args[1] == "open" && os.Args[2] == "--today":
		today = true
	default:
		fmt.Fprintln(os.Stderr, "usage: textedit open [--today]")
		os.Exit(2)
	}

	if err := open(today); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// Minimal RTF document; TextEdit can't open a completely empty .rtf file.
const starterContent = `{\rtf1\ansi\ansicpg1252\cocoartf2907
\cocoatextscaling0\cocoaplatform0{\fonttbl}
{\colortbl;\red255\green255\blue255;}
{\*\expandedcolortbl;;}
\margl1440\margr1440\vieww11520\viewh8400\viewkind0
}
`

func createIfMissing(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if errors.Is(err, fs.ErrExist) {
		return nil
	}
	if err != nil {
		return err
	}
	if _, err := f.WriteString(starterContent); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}

func open(today bool) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	base := filepath.Join(home, "textedit")
	path := filepath.Join(base, fileName)
	if today {
		dir := filepath.Join(base, time.Now().Format("2006-01-02"))
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
		path = filepath.Join(dir, fileName)
		if err := createIfMissing(path); err != nil {
			return err
		}
	} else if _, err := os.Stat(path); err != nil {
		return err
	}

	if out, err := exec.Command("open", "-a", "TextEdit", path).CombinedOutput(); err != nil {
		return fmt.Errorf("opening %s: %w\n%s", path, err, out)
	}
	return nil
}
