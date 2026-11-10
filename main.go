package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// The file name is a single braille blank character (U+2800) followed by ".rtf".
const fileName = "\u2800.rtf"

func main() {
	if len(os.Args) != 2 || os.Args[1] != "open" {
		fmt.Fprintln(os.Stderr, "usage: textedit open")
		os.Exit(2)
	}

	if err := open(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func open() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	path := filepath.Join(home, "textedit", fileName)
	if _, err := os.Stat(path); err != nil {
		return err
	}

	if out, err := exec.Command("open", "-a", "TextEdit", path).CombinedOutput(); err != nil {
		return fmt.Errorf("opening %s: %w\n%s", path, err, out)
	}
	return nil
}
