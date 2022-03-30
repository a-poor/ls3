package lib

import (
	"io/fs"
)

type FileObject struct {
	Path  string // Path where the file/dir is located
	Name  string // Name of the file/dir
	IsDir bool   // Is this a directory?
	Size  int64  // Size of the file in bytes
}

func newFileObject(dir string, fi fs.FileInfo) FileObject {
	if fi.IsDir() {
		return newDirFO(dir, fi.Name())
	}
	return newFileFO(dir, fi.Name(), fi.Size())
}

func newFileFO(dir, name string, size int64) FileObject {
	return FileObject{
		Path:  dir,
		Name:  name,
		IsDir: false,
		Size:  size,
	}
}

func newDirFO(dir, name string) FileObject {
	return FileObject{
		Path:  dir,
		Name:  name,
		IsDir: true,
		Size:  0,
	}
}

type FileSystem interface {
	// List the contents of the current working directory
	ListContents() ([]FileObject, error)

	// Change the current working directory
	ChangeDir(string) error

	// Read the contents of a file
	GetFile(string) ([]byte, error)

	// Write the contents of a file
	WriteFile(string, []byte) error

	// DoesDirExist(string) (bool, error)
	// DoesFileExist(string) (bool, error)
}
