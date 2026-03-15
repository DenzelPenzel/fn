package fn

import "os"

// WriteFile writes data to a temporary file and atomically renames it to the
// target path, ensuring the file is fully written before it appears at the
// destination
func WriteFile(path string, data []byte, perm os.FileMode) error {
	tmp := path + ".tmp"

	err := os.WriteFile(tmp, data, perm)
	if err != nil {
		return err
	}

	return os.Rename(tmp, path)
}

// WriteFileRemove is like WriteFile but removes the destination first if it
// exists
func WriteFileRemove(path string, data []byte, perm os.FileMode) error {
	_ = os.Remove(path)
	return WriteFile(path, data, perm)
}
