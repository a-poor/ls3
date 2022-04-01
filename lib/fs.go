package lib

import (
	"errors"
	"io/fs"
	"sort"
)

var (
	ErrFileDoesNotExist  = errors.New("file does not exist")
	ErrDirDoesNotExist   = errors.New("directory does not exist")
	ErrPathAlreadyExists = errors.New("path already exists")
	ErrExpectedDirPath   = errors.New("expected path to be a directory")
	ErrExpectedFilePath  = errors.New("expected path to be a file")
)

// FileObject represents a file or directory in the filesystem.
type FileObject struct {
	Name  string // Name of the file/dir
	IsDir bool   // Is this a directory?
	Size  int64  // Size of the file in bytes
}

// newFileObject creates a new FileObject instances from the given fs.FileInfo.
func newFileObject(fi fs.FileInfo) FileObject {
	if fi.IsDir() {
		return newDirFO(fi.Name())
	}
	return newFileFO(fi.Name(), fi.Size())
}

// newFileFO creates a new FileObject instance for a file.
func newFileFO(name string, size int64) FileObject {
	return FileObject{
		Name:  name,
		IsDir: false,
		Size:  size,
	}
}

// newDirFO creates a new FileObject instance for a directory.
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

// FileSystem is an interface for standardizing the process
// of navigating directories and accessing files in a local
// filesystem or in S3.
//
// The FileSystem interface is used by the LocalFS, for
// accessing local files (or any billy.Filesystem, like Memfs),
// and S3FS, for accessing files in S3.
//
// Currently, FileSystem only supports listing a directory's
// contents, changing directories, getting files, and writing
// files.
//
// Future versions of this interface may support other operations
// such as creating directories, deleting files, copying files from
// another FileSystem checking if a file/dir exists, previewing
// files, diffing files, etc.
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
	// CopyFileFrom(string, string, FileSystem) error
}
