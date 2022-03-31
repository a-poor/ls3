package lib_test

import (
	"testing"

	"github.com/a-poor/ls3/lib"
)

func TestFileObject(t *testing.T) {
	// Define some files and directories
	d1 := lib.FileObject{
		Name:  "_dir",
		IsDir: true,
	}
	d2 := lib.FileObject{
		Name:  "foo",
		IsDir: true,
	}
	f1 := lib.FileObject{
		Name: "bar.txt",
	}
	f2 := lib.FileObject{
		Name: ".baz.json",
	}

	// Add them to a list
	fos := []lib.FileObject{f1, d2, d1, f2}
	lib.SortFileObjects(fos)

	// Check the list
	expectations := []string{
		"_dir",
		"foo",
		".baz.json",
		"bar.txt",
	}
	for i, exp := range expectations {
		if fos[i].Name != exp {
			t.Errorf("Expected fos[%d] to be %q, got %q", i, exp, fos[0].Name)
		}
	}

}
