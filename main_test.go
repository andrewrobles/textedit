package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestAppendFile(t *testing.T) {
	source := filepath.Join("test", "source.rtf")
	expected := filepath.Join("test", "expected.rtf")

	// Work on a copy so the destination.rtf fixture is left untouched.
	destination := filepath.Join(t.TempDir(), "destination.rtf")
	fixture, err := os.ReadFile(filepath.Join("test", "destination.rtf"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(destination, fixture, 0o644); err != nil {
		t.Fatal(err)
	}

	if err := appendFile(source, destination); err != nil {
		t.Fatalf("appendFile(%q, %q) returned error: %v", source, destination, err)
	}

	got, err := os.ReadFile(destination)
	if err != nil {
		t.Fatal(err)
	}
	want, err := os.ReadFile(expected)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("destination contents do not match expected.rtf\ngot:\n%s\nwant:\n%s", got, want)
	}
}
