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
