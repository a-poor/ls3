package lib

import (
	"errors"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
)

const (
	// Root directory of the billy filesystem
	RootBillyDir = "/" // TODO – Is this the best option? Can this be reconfigured?
)

// LocalFS implements the FileSystem interface for the local filesystem.
// It uses a billy.Filesystem for accessing the local fs to allow for testing
// with a memfs.
type LocalFS struct {
	FS      billy.Filesystem // Underlying filesystem (e.g. osfs or memfs)
	WorkDir string           // Current working directory
}

// NewLocalFS creates a new LocalFS instance with the given baseDir as the
// starting working directory.
func NewLocalFS(baseDir string) *LocalFS {
	fs := osfs.New(RootBillyDir)
	return NewLocalFSFromBillyFS(fs, baseDir)
}

// NewLocalFSFromBillyFS creates a new LocalFS instance based on the given
// billy.Filesystem and with baseDir as the starting working directory.
//
// This is useful for testing, where you can create a memfs and prepopulate
// it with files, and then pass it to this function to create a LocalFS.
func NewLocalFSFromBillyFS(fs billy.Filesystem, baseDir string) *LocalFS {
	return &LocalFS{
		FS:      fs,
		WorkDir: baseDir,
	}
}

func (fs *LocalFS) ListContents() ([]FileObject, error) {
	// Get the current working directory's contents
	fis, err := fs.FS.ReadDir(fs.WorkDir)
	if err != nil {
		return nil, err
	}

	// Convert them to FileObjects
	var files []FileObject
	for _, fi := range fis {
		files = append(files, newFileObject(fi))
	}

	// Sort the files and return
	SortFileObjects(files)
	return files, nil
}

func (fs *LocalFS) ChangeDir(newDir string) error {
	return errors.New("not implemented")
}

func (fs *LocalFS) GetFile(fileName string) ([]byte, error) {
	return nil, errors.New("not implemented")
}

func (fs *LocalFS) WriteFile(fileName string, b []byte) error {
	return errors.New("not implemented")
}
