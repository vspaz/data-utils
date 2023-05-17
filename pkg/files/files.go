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
		log.Printf("error reading a file: '%s'\n", err)
	}
	return fh
}
