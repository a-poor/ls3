package lib_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/a-poor/ls3/lib"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
)

// splitPath is a helper function for getting the first
// non-empty segment of the path p.
func splitPath(p string) string {
	// Check if the path is empty
	if p == "" {
		return ""
	}

	// Split the path
	sep := "/"
	parts := strings.Split(p, sep)

	for _, p := range parts {
		if p != "" {
			return p
		}
	}
	return "" // Error: Return empty string
}

func fmtFileBody(name string) string {
	return fmt.Sprintf("This file is called %q\n", name)
}

func prepareMemFS() (billy.Filesystem, error) {
	// Define files and directories to create
	dirs := []string{
		"/.git",
		"/_dir2",
		"/dir3/subdir",
	}
	files := []string{
		"/_foo.txt",
		"/bar.json",
		"/.gitignore",
		"/_dir2/sub-file.txt",
		"/dir3/subdir/another_file",
	}

	// Prepare a billy filesystem to test with
	fs := memfs.New()
	for _, d := range dirs {
		err := fs.MkdirAll(d, 0755)
		if err != nil {
			return nil, err
		}
	}
	for _, f := range files {
		file, err := fs.Create(f)
		if err != nil {
			return nil, err
		}
		txt := fmtFileBody(f)
		file.Write([]byte(txt))
		err = file.Close()
		if err != nil {
			return nil, err
		}
	}
	// ...done with billly FS prep.
	return fs, nil
}

func TestLocalFSListContents(t *testing.T) {
	fs, err := prepareMemFS()
	if err != nil {
		t.Errorf("Error preparing billy.Filesystem: %s", err)
		t.FailNow()
	}

	// Create a new LocalFS instance with the billy filesystem
	lfs := lib.NewLocalFSFromBillyFS(fs, "/")
	contents, err := lfs.ListContents()
	if err != nil {
		t.Errorf("Error listing contents: %s", err)
		t.FailNow()
	}

	// Check the directories...
	expectDirs := []string{
		"/.git",
		"/_dir2",
		"/dir3",
	}
	for _, d := range expectDirs {
		p := splitPath(d)
		found := false
		for _, c := range contents {
			if c.Name != p {
				continue
			}
			found = true
			if !c.IsDir {
				t.Errorf("Expected %q to be a directory", d)
			}
			break
		}
		if !found {
			t.Errorf("Directory %q (%q) not found in LocalFS's workdir", p, d)
		}
	}

	// Check the files...
	expectFiles := []string{
		"/_foo.txt",
		"/bar.json",
		"/.gitignore",
	}
	for _, f := range expectFiles {
		p := splitPath(f)
		found := false
		for _, c := range contents {
			if c.Name != p {
				continue
			}
			found = true
			if c.IsDir {
				t.Errorf("Expected %q (%q) to be a file", f, p)
			}
			break
		}
		if !found {
			t.Errorf("File %q (%q) not found in LocalFS's workdir", p, f)
		}
	}
}

func TestLocalFSChangeDir(t *testing.T) {
	fs, err := prepareMemFS()
	if err != nil {
		t.Errorf("Error preparing billy.Filesystem: %s", err)
		t.FailNow()
	}

	// Create a new LocalFS instance with the billy filesystem
	lfs := lib.NewLocalFSFromBillyFS(fs, "/")

	exp := "/"
	if lfs.WorkDir != exp {
		t.Errorf("Expected %q, got %q", exp, lfs.WorkDir)
	}

	// Move into the first directory
	err = lfs.ChangeDir("dir1")
	if err != nil {
		t.Errorf("Error changing dir: %s", err)
		t.FailNow()
	}

	exp = "/dir1"
	if lfs.WorkDir != exp {
		t.Errorf("Expected %q, got %q", exp, lfs.WorkDir)
		t.FailNow()
	}

	// Move back up again
	err = lfs.ChangeDir("..")
	if err != nil {
		t.Errorf("Error changing dir: %s", err)
		t.FailNow()
	}

	exp = "/"
	if lfs.WorkDir != exp {
		t.Errorf("Expected %q, got %q", exp, lfs.WorkDir)
		t.FailNow()
	}

	// Can't move up any further
	err = lfs.ChangeDir("..")
	if err != nil {
		t.Errorf("Error changing dir: %s", err)
		t.FailNow()
	}

	exp = "/"
	if lfs.WorkDir != exp {
		t.Errorf("Expected %q, got %q", exp, lfs.WorkDir)
		t.FailNow()
	}

	// Move down into dir3
	err = lfs.ChangeDir("dir3")
	if err != nil {
		t.Errorf("Error changing dir: %s", err)
		t.FailNow()
	}

	exp = "/dir3"
	if lfs.WorkDir != exp {
		t.Errorf("Expected %q, got %q", exp, lfs.WorkDir)
		t.FailNow()
	}

	// Move down to the subdir
	err = lfs.ChangeDir("subdir")
	if err != nil {
		t.Errorf("Error changing dir: %s", err)
		t.FailNow()
	}

	exp = "/dir3/subdir"
	if lfs.WorkDir != exp {
		t.Errorf("Expected %q, got %q", exp, lfs.WorkDir)
		t.FailNow()
	}

	// TODO - Add test for changing to a non-existent directory
	// TODO - Add test for changing to a non-directory (aka file)
}

func TestLocalFSGetFile(t *testing.T) {
	// Create a new pre-populated billy fs and LocalFS instance
	fs, err := prepareMemFS()
	if err != nil {
		t.Errorf("Error preparing billy.Filesystem: %s", err)
		t.FailNow()
	}
	lfs := lib.NewLocalFSFromBillyFS(fs, "/")

	fileNames := []string{
		"/_foo.txt",
		"/bar.json",
	}

	for _, fn := range fileNames {
		b, err := lfs.GetFile(fn)
		if err != nil {
			t.Errorf("Error getting file %q: %s", fn, err)
			t.FailNow()
		}

		expect := fmtFileBody(fn)
		body := string(b)
		if body != expect {
			t.Errorf("Expected %q, got %q", expect, body)
		}
	}
}

func TestLocalFSWriteFile(t *testing.T) {
	// Create a new pre-populated billy fs and LocalFS instance
	fs, err := prepareMemFS()
	if err != nil {
		t.Errorf("Error preparing billy.Filesystem: %s", err)
		t.FailNow()
	}
	// lfs := lib.NewLocalFSFromBillyFS(fs, "/")
	_ = lib.NewLocalFSFromBillyFS(fs, "/")
}
