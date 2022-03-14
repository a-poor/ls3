package main

import "os"

// Use billy to read local files

type localFile struct {
	title string
	isDir bool
}

func getWorkingDir() localFile {
	return localFile{".", true}
}

func (f localFile) Title() string {
	n := f.title
	if f.isDir {
		n += "/"
	}
	return n
}

func (f localFile) Description() string {
	return "<more-info>"
}

func (f localFile) FilterValue() string {
	return f.Title()
}

func getLocalFiles() ([]localFile, error) {
	files, err := os.ReadDir(".")
	if err != nil {
		return nil, err
	}
	fs := []localFile{getWorkingDir()}
	for _, f := range files {
		fs = append(fs, localFile{f.Name(), f.IsDir()})
	}
	return fs, nil
}
