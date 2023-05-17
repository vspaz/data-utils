package files

func MustClose(fh *os.File) {
	if err := fh.Close(); err != nil {
		log.Fatal(err)
	}
}
