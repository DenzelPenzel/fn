package fn

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")

	err := WriteFile(path, []byte("hello"), 0644)
	assertNoError(t, err)

	data, err := os.ReadFile(path)
	assertNoError(t, err)
	assertEqual(t, string(data), "hello")

	// Tmp file should not exist
	_, err = os.Stat(path + ".tmp")
	assertTrue(t, os.IsNotExist(err))
}

func TestWriteFileRemove(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")

	// Write initial content
	err := os.WriteFile(path, []byte("old"), 0600)
	assertNoError(t, err)

	// Overwrite with WriteFileRemove
	err = WriteFileRemove(path, []byte("new"), 0644)
	assertNoError(t, err)

	data, err := os.ReadFile(path)
	assertNoError(t, err)
	assertEqual(t, string(data), "new")
}
