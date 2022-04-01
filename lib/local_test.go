package lib_test

import (
	"fmt"
	"io/ioutil"
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

// fmtFileBody is a helper function for creating a consistent
// body for test files.
func fmtFileBody(name string) string {
	return fmt.Sprintf("This file is called %q\n", name)
}

// prepareMemFS is a helper function for creating and populating
// an in-memory billy.Filesystem.
func prepareMemFS() (billy.Filesystem, error) {
	// Define files and directories to create
	dirs := []string{
		"/.git",
		"/dir1",
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
		_, err = file.Write([]byte(txt))
		if err != nil {
			return nil, err
		}

		err = file.Close()
		if err != nil {
			return nil, err
		}
	}
	// ...done with billly FS prep.
	return fs, nil
}

func Test_prepareMemFS(t *testing.T) {
	fs, err := prepareMemFS()
	if err != nil {
		t.Errorf("Error preparing billy.Filesystem: %s", err)
		t.FailNow()
	}

	// Check the contents of the billy filesystem
	files := []string{
		"/_foo.txt",
		"/bar.json",
		"/.gitignore",
		"/_dir2/sub-file.txt",
		"/dir3/subdir/another_file",
	}
	for _, fn := range files {
		f, err := fs.Open(fn)
		if err != nil {
			t.Errorf("Error opening file %q: %s", fn, err)
			t.FailNow()
		}
		defer f.Close()

		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Errorf("Error reading file %q: %s", fn, err)
			t.FailNow()
		}

		txt := fmtFileBody(fn)
		if string(b) != txt {
			t.Errorf("Expected %q, got %q", txt, string(b))
		}
	}
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
			t.Errorf("Expected file %q's contents to be %q, got %q", fn, expect, body)
		}
	}

	// TODO - Add test for getting a non-existent file
}

func TestLocalFSWriteFile(t *testing.T) {
	// Create a new pre-populated billy fs and LocalFS instance
	fs, err := prepareMemFS()
	if err != nil {
		t.Errorf("Error preparing billy.Filesystem: %s", err)
		t.FailNow()
	}
	lfs := lib.NewLocalFSFromBillyFS(fs, "/")

	// TODO - Add test for writing a file
	files := []string{
		"foo.txt",
		"bar.txt",
		"baz.txt",
	}
	for _, fn := range files {
		txt := fmtFileBody(fn)
		err = lfs.WriteFile(fn, []byte(txt))
		if err != nil {
			t.Errorf("Error writing file to LocalFS %q: %s", fn, err)
			continue
		}
		f, err := fs.Open(fn)
		if err != nil {
			t.Errorf("Error opening billy.File %q: %s", fn, err)
			continue
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Errorf("Error reading billy.File %q: %s", fn, err)
			continue
		}
		if string(b) != txt {
			t.Errorf("Expected file %q to be %q, got %q", fn, txt, string(b))
			continue
		}
	}

	// TODO - Add test for writing a file to a different directory...
}

func TestPathExists(t *testing.T) {
	// Create a new pre-populated billy fs and LocalFS instance
	fs, err := prepareMemFS()
	if err != nil {
		t.Errorf("Error preparing billy.Filesystem: %s", err)
		t.FailNow()
	}
	lfs := lib.NewLocalFSFromBillyFS(fs, "/")

	files := []string{
		"/_foo.txt",
		"/bar.json",
		"/.gitignore",
	}
	dirs := []string{
		".git/",
		"_dir2",
		"dir3/",
	}
	fakePaths := []string{
		".not-really-git/",
		"dirrrrr/",
		"fake.txt",
		"im-not-here.json",
	}
	t.Run("files-exist", func(t *testing.T) {
		for _, fn := range files {
			if !lfs.PathExists(fn) {
				t.Errorf("Expected file path %q to exist", fn)
			}
		}
	})
	t.Run("dirs-exist", func(t *testing.T) {
		for _, dn := range dirs {
			if !lfs.PathExists(dn) {
				t.Errorf("Expected dir path %q to exist", dn)
			}
		}
	})
	t.Run("fake-paths", func(t *testing.T) {
		for _, p := range fakePaths {
			if lfs.PathExists(p) {
				t.Errorf("Expected fake path %q to not exist", p)
			}
		}
	})

	// TODO - Test checking paths after changing directories
}

func TestIsFile(t *testing.T) {
	// Create a new pre-populated billy fs and LocalFS instance
	fs, err := prepareMemFS()
	if err != nil {
		t.Errorf("Error preparing billy.Filesystem: %s", err)
		t.FailNow()
	}
	lfs := lib.NewLocalFSFromBillyFS(fs, "/")

	files := []string{
		"/_foo.txt",
		"/bar.json",
		"/.gitignore",
	}
	dirs := []string{
		".git/",
		"_dir2",
		"dir3/",
	}
	fakePaths := []string{
		".not-really-git/",
		"dirrrrr/",
		"fake.txt",
		"im-not-here.json",
	}
	t.Run("files-exist", func(t *testing.T) {
		for _, fn := range files {
			if !lfs.IsFile(fn) {
				t.Errorf("Expected file path %q to be file", fn)
			}
		}
	})
	t.Run("dirs-not-files", func(t *testing.T) {
		for _, dn := range dirs {
			if lfs.IsFile(dn) {
				t.Errorf("Expected dir path %q to not be file", dn)
			}
		}
	})
	t.Run("fake-paths", func(t *testing.T) {
		for _, p := range fakePaths {
			if lfs.IsFile(p) {
				t.Errorf("Expected fake path %q to not be file", p)
			}
		}
	})

	// TODO - Test checking paths after changing directories
}

func TestIsDir(t *testing.T) {
	// Create a new pre-populated billy fs and LocalFS instance
	fs, err := prepareMemFS()
	if err != nil {
		t.Errorf("Error preparing billy.Filesystem: %s", err)
		t.FailNow()
	}
	lfs := lib.NewLocalFSFromBillyFS(fs, "/")

	files := []string{
		"/_foo.txt",
		"/bar.json",
		"/.gitignore",
	}
	dirs := []string{
		".git/",
		"_dir2",
		"dir3/",
	}
	fakePaths := []string{
		".not-really-git/",
		"dirrrrr/",
		"fake.txt",
		"im-not-here.json",
	}
	t.Run("files-not-dirs", func(t *testing.T) {
		for _, fn := range files {
			if lfs.IsDir(fn) {
				t.Errorf("Expected file path %q to not be dir", fn)
			}
		}
	})
	t.Run("dirs-exist", func(t *testing.T) {
		for _, dn := range dirs {
			if !lfs.IsDir(dn) {
				t.Errorf("Expected dir path %q to be dir", dn)
			}
		}
	})
	t.Run("fake-paths", func(t *testing.T) {
		for _, p := range fakePaths {
			if lfs.IsDir(p) {
				t.Errorf("Expected fake path %q to not be dir", p)
			}
		}
	})

	// TODO - Test checking paths after changing directories
}
