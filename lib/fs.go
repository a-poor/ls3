package lib

import (
	"io/fs"
	"sort"
)

// FileObject represents a file or directory in the filesystem.
type FileObject struct {
	Name  string // Name of the file/dir
	IsDir bool   // Is this a directory?
	Size  int64  // Size of the file in bytes
}

func newFileObject(fi fs.FileInfo) FileObject {
	if fi.IsDir() {
		return newDirFO(fi.Name())
	}
	return newFileFO(fi.Name(), fi.Size())
}

func newFileFO(name string, size int64) FileObject {
	return FileObject{
		Name:  name,
		IsDir: false,
		Size:  size,
	}
}

func newDirFO(name string) FileObject {
	return FileObject{
		Name:  name,
		IsDir: true,
		Size:  0,
	}
}

// SortFileObjects sorts the given slice of FileObjects in place.
func SortFileObjects(fos []FileObject) {
	sort.Slice(fos, func(i, j int) bool {
		if fos[i].IsDir == fos[j].IsDir {
			return fos[i].Name < fos[j].Name
		}
		return fos[i].IsDir
	})
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
