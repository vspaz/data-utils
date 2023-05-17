package files

import (
	"log"
	"os"
)

func MustClose(fh *os.File) {
	if err := fh.Close(); err != nil {
		log.Fatal(err)
	}
}

func FromFile(path string) *os.File {
	log.Printf("trying to read file at '%s'\n", path)
	fh, err := os.Open(path)
	if err != nil {
		log.Fatalf("error reading a file: '%s'\n", err)
	}
	return fh
}

func CreateFile(path string) *os.File {
	fh, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed to create a file: '%s'\n", err)
	}
	return fh
}
