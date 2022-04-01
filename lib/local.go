package lib

import (
	"io/ioutil"
	"os"
	"path"

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

func (fs *LocalFS) fmtPath(p string) string {
	p2 := path.Join(fs.WorkDir, p)
	p2 = path.Clean(p2)
	return p2
}

func (fs *LocalFS) ChangeDir(newDir string) error {
	// Validate the new path
	if !fs.PathExists(newDir) {
		return ErrDirDoesNotExist
	}
	if !fs.IsDir(newDir) {
		return ErrExpectedDirPath
	}

	// Format the new path
	wd := fs.fmtPath(newDir)

	// Set the new path and return success
	fs.WorkDir = wd
	return nil
}

func (fs *LocalFS) GetFile(fileName string) ([]byte, error) {
	// Validate the path
	if !fs.PathExists(fileName) {
		return nil, ErrFileDoesNotExist
	}
	if !fs.IsFile(fileName) {
		return nil, ErrExpectedFilePath
	}

	// Format the new path
	wd := fs.fmtPath(fileName)

	// Open & read the file
	f, err := fs.FS.Open(wd)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (fs *LocalFS) WriteFile(fileName string, b []byte) error {
	// Validate the path
	if fs.IsDir(fileName) {
		return ErrExpectedFilePath
	}

	// Format the new path
	wd := fs.fmtPath(fileName)

	// Get the file
	f, err := fs.FS.Create(wd)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the file
	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func (fs *LocalFS) PathExists(name string) bool {
	p := fs.fmtPath(name)
	_, err := fs.FS.Stat(p)
	return !os.IsNotExist(err)
}

func (fs *LocalFS) IsFile(name string) bool {
	p := fs.fmtPath(name)
	fi, _ := fs.FS.Stat(p) // TODO - Handle error?
	return fi != nil && !fi.IsDir()
}

func (fs *LocalFS) IsDir(name string) bool {
	p := fs.fmtPath(name)
	fi, _ := fs.FS.Stat(p) // TODO - Handle error?
	return fi != nil && fi.IsDir()
}
