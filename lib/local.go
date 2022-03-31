package lib

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
)

type LocalFS struct {
	FS       billy.Filesystem
	WorkPath string
}

func NewLocalFS(baseDir string) *LocalFS {
	fs := osfs.New(baseDir)
	return NewLocalFSFromBillyFS(fs)
}

func NewLocalFSFromBillyFS(fs billy.Filesystem) *LocalFS {
	return &LocalFS{
		FS:       fs,
		WorkPath: fs.Root(),
	}
}

func (fs *LocalFS) ListContents() ([]FileObject, error) {
	fis, err := fs.FS.ReadDir(fs.WorkPath)
	if err != nil {
		return nil, err
	}

	var files []FileObject
	for _, fi := range fis {
		files = append(files, newFileObject(fi))
	}
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
