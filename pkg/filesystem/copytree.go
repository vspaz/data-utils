package filesystem

import (
	"io"
	"os"
	"path/filepath"
)

// copytree provides similar functionality to Python shutil.copytree
// https://docs.python.org/3.11/library/shutil.html#shutil.copytree

// FileSystem FileSystem implements a FileSystem interface.
type FileSystem struct{}

// New returns a FileSystem object.
func New() *FileSystem {
	return &FileSystem{}
}

// copyFile copes contents of the file named src to a file named dst.
func (fs *FileSystem) CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		MustClose(out)
		MustClose(in)
	}()
	if _, err = io.Copy(out, in); err != nil {
		return err
	}
	err = out.Sync()
	return nil
}

// CopyTree recursively creates and populates directory tree.
// similar to python shutil.copytree
func (fs *FileSystem) CopyTree(src, dst string) error {
	files, err := os.ReadDir(src)
	if err != nil {
		return err
	}
	if err = os.Mkdir(dst, os.ModePerm); err != nil {
		return err
	}
	for _, file := range files {
		srcName := filepath.Join(src, file.Name())
		dstName := filepath.Join(dst, file.Name())
		fileInfo, err := os.Stat(srcName)
		if err != nil {
			return err
		}

		if fileInfo.IsDir() {
			if err = fs.CopyTree(srcName, dstName); err != nil {
				return err
			}
		} else {
			if err = fs.CopyFile(srcName, dstName); err != nil {
				return err
			}
		}
	}
	return nil
}
