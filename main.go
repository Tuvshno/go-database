package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println()
}

func saveDataNaive(path string, data []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0665)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(data)
	return err

}

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
