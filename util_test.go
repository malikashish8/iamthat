package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolvePathUserFolder(t *testing.T) {
	resolvedPath := ResolvePath("~/foo.txt")

	userHomeDir, _ := os.UserHomeDir()
	absolutePath := filepath.Join(userHomeDir, "foo.txt")

	if absolutePath != resolvedPath {
		t.Errorf("ResolvePath(%q) == %q, want %q", "~/foo.txt", resolvedPath, absolutePath)
	}
}

func TestResolvePathCurrentFolder(t *testing.T) {
	resolvedPath := ResolvePath("./foo.txt")

	wd, _ := os.Getwd()
	absolutePath := filepath.Join(wd, "foo.txt")

	if absolutePath != resolvedPath {
		t.Errorf("ResolvePath(%q) == %q, want %q", "./foo.txt", resolvedPath, absolutePath)
	}
}

func TestCheckFilesSameContent(t *testing.T) {
	contents := `hi:\nhello:\nid: 123`
	f1, _ := os.CreateTemp("", "f1")
	f2, _ := os.CreateTemp("", "f2")

	defer os.Remove(f1.Name())
	defer os.Remove(f2.Name())

	f1.Write([]byte(contents))
	f2.Write([]byte(contents))

	if !CheckFilesSameContent(f1.Name(), f2.Name()) {
		t.Errorf("CheckFilesSameContent returned false for same files")
	}
}

func TestCheckFilesSameContent2(t *testing.T) {
	content1 := `hi:\nhello:\nid: 123`
	content2 := `abc:\ndef: 1123`
	f1, _ := os.CreateTemp("", "f1")
	f2, _ := os.CreateTemp("", "f2")

	defer os.Remove(f1.Name())
	defer os.Remove(f2.Name())

	f1.Write([]byte(content1))
	f2.Write([]byte(content2))

	if CheckFilesSameContent(f1.Name(), f2.Name()) {
		t.Errorf("CheckFilesSameContent returned true for differrent files")
	}
}
