package lib

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
)

type LocalFS struct {
	FS      billy.Filesystem
	WorkDir string
}

func NewLocalFS(baseDir string) *LocalFS {
	fs := osfs.New("/")
	return NewLocalFSFromBillyFS(fs, baseDir)
}

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
	return nil
}

func (fs *LocalFS) GetFile(fileName string) ([]byte, error) {
	return nil, nil
}

func (fs *LocalFS) WriteFile(fileName string, b []byte) error {
	return nil
}
