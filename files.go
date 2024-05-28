// This shows the problems with saving data to files rather than databases

package main

import (
	"fmt"
	"os"
)

// This is the naive approach to storing data to a file
// Opens file in write only, Creates the file if not there, Truncates (overwrites) writable file when opened, Then writes the data to the file
// Problems: Trunactes the file which prevents concurrency. Writing is not atomic, data may not be persisted to disk
func saveDataNaive(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0665)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	return err

}

// This is a better approach by dumping the data into a temp file and then renaming
// Renaming is an atomic option
// Problems: Doesn't control when the data is persisted to the disk because metadata may be saved before the data, causing curroption on crash
func saveDataBetter(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0604)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}

	return os.Rename(tmp, path)
}

// Now we have flushed the data to the sync
// Problem: we also have to save the metadata. This is the main problem with saving data with files
func saveDataBetterFSync(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, randomInt())
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0604)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	if err != nil {
		os.Remove(tmp)
		return err
	}

	err = fp.Sync()
	if err != nil {
		os.Remove(tmp)
		return err
	}

	return os.Rename(tmp, path)
}
