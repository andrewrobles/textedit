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
	var err error
	switch {
	case len(os.Args) == 2 && os.Args[1] == "open":
		err = open(false)
	case len(os.Args) == 3 && os.Args[1] == "open" && os.Args[2] == "--today":
		err = open(true)
	case len(os.Args) == 2 && os.Args[1] == "commit":
		err = commit()
	case len(os.Args) == 2 && os.Args[1] == "save":
		err = save()
	default:
		fmt.Fprintln(os.Stderr, "usage: textedit open [--today] | textedit commit | textedit save")
		os.Exit(2)
	}

	if err != nil {
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

func baseDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "textedit"), nil
}

func open(today bool) error {
	base, err := baseDir()
	if err != nil {
		return err
	}

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

	return openInTextEdit(path)
}

// save archives the source RTF file to a timestamped copy, resets the source to the
// starter content, commits, and reloads the source in TextEdit.
func save() error {
	base, err := baseDir()
	if err != nil {
		return err
	}

	source := filepath.Join(base, fileName)
	data, err := os.ReadFile(source)
	if err != nil {
		return err
	}
	// Basic ISO 8601 format: colons are awkward in macOS file names.
	archive := filepath.Join(base, time.Now().Format("20060102T150405")+".rtf")
	if err := os.WriteFile(archive, data, 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(source, []byte(starterContent), 0o644); err != nil {
		return err
	}

	// Close without saving so TextEdit doesn't write its stale contents back.
	if err := closeInTextEdit(source); err != nil {
		return err
	}
	if err := commit(); err != nil {
		return err
	}
	return openInTextEdit(source)
}

const closeScript = `on run argv
	set p to item 1 of argv
	if application "TextEdit" is running then
		tell application "TextEdit" to close (every document whose path is p) saving no
	end if
end run`

func closeInTextEdit(path string) error {
	if out, err := exec.Command("osascript", "-e", closeScript, path).CombinedOutput(); err != nil {
		return fmt.Errorf("closing %s: %w\n%s", path, err, out)
	}
	return nil
}

func openInTextEdit(path string) error {
	if out, err := exec.Command("open", "-a", "TextEdit", path).CombinedOutput(); err != nil {
		return fmt.Errorf("opening %s: %w\n%s", path, err, out)
	}
	return nil
}

// commit stages all changes in the textedit directory and commits them with an empty message.
func commit() error {
	base, err := baseDir()
	if err != nil {
		return err
	}

	if err := git(base, "add", "--all"); err != nil {
		return err
	}
	return git(base, "commit", "--allow-empty-message", "--message", "")
}

func git(dir string, args ...string) error {
	cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git %s: %w", args[0], err)
	}
	return nil
}
