package main

import (
	"fmt"
	"log"
	"os"
)

func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp.%d", path, 500)
	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)
	if err != nil {
		return err
	}
	defer func() {
		fp.Close()
		if err != nil {
			os.Remove(tmp)
		}
	}()
	_, err = fp.Write(data)
	if err != nil {
		return err
	}
	err = fp.Sync() // fsync
	if err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func main() {
	content := []byte("This is a test of the atomic save function2.")
	filePath := "output.txt"

	fmt.Println("Attempting to save data to", filePath)
	if err := SaveData2(filePath, content); err != nil {
		log.Fatalf("ERROR: Could not save the file. Reason: %v", err)
	}
	fmt.Println("Success! The file was saved.")
}
