package lib_test

import (
	"testing"
)

func TestS3FSListContents(t *testing.T) {
	// // Define files and directories to create
	// dirs := []string{
	// 	"/.git",
	// 	"/_dir2",
	// 	"/dir3/subdir",
	// }
	// files := []string{
	// 	"/_foo.txt",
	// 	"/bar.json",
	// 	"/.gitignore",
	// 	"/_dir2/sub-file.txt",
	// 	"/dir3/subdir/another_file",
	// }

	// // Prepare a billy filesystem to test with
	// fs := memfs.New()
	// for _, d := range dirs {
	// 	err := fs.MkdirAll(d, 0755)
	// 	if err != nil {
	// 		t.Errorf("Error preparing billy.Filesystem dirs: %s", err)
	// 		t.FailNow()
	// 	}
	// }
	// for _, f := range files {
	// 	_, err := fs.Create(f)
	// 	if err != nil {
	// 		t.Errorf("Error preparing billy.Filesystem files: %s", err)
	// 		t.FailNow()
	// 	}
	// }
	// // ...done with billly FS prep.

	// // Create a new S3FS instance with the billy filesystem
	// lfs := lib.NewS3FSFromBillyFS(fs)
	// contents, err := lfs.ListContents()
	// if err != nil {
	// 	t.Errorf("Error listing contents: %s", err)
	// 	t.FailNow()
	// }

	// // Check the directories...
	// expectDirs := []string{
	// 	"/.git",
	// 	"/_dir2",
	// 	"/dir3",
	// }
	// for _, d := range expectDirs {
	// 	p := splitPath(d)
	// 	found := false
	// 	for _, c := range contents {
	// 		if c.Name != p {
	// 			continue
	// 		}
	// 		found = true
	// 		if !c.IsDir {
	// 			t.Errorf("Expected %q to be a directory", d)
	// 		}
	// 		break
	// 	}
	// 	if !found {
	// 		t.Errorf("Directory %q (%q) not found in S3FS's workdir", p, d)
	// 	}
	// }

	// // Check the files...
	// expectFiles := []string{
	// 	"/_foo.txt",
	// 	"/bar.json",
	// 	"/.gitignore",
	// }
	// for _, f := range expectFiles {
	// 	p := splitPath(f)
	// 	found := false
	// 	for _, c := range contents {
	// 		if c.Name != p {
	// 			continue
	// 		}
	// 		found = true
	// 		if c.IsDir {
	// 			t.Errorf("Expected %q (%q) to be a file", f, p)
	// 		}
	// 		break
	// 	}
	// 	if !found {
	// 		t.Errorf("File %q (%q) not found in S3FS's workdir", p, f)
	// 	}
	// }
}

func TestS3FSChangeDir(t *testing.T) {

}

func TestS3FSGetFile(t *testing.T) {

}

func TestS3FSWriteFile(t *testing.T) {

}
