package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	startPath := "."
	rootFolder := newFolder(startPath)

	visit := func(path string, info os.FileInfo, err error) error {

		segments := strings.Split(path, string(filepath.Separator))
		if info.IsDir() {
			if path != startPath {
				rootFolder.addFolder(segments)
			}
		} else {
			rootFolder.addFile(segments)
		}
		return nil
	}

	err := filepath.Walk(".", visit)
	if err != nil {
		log.Fatal(err)
	}

	for _, folder := range rootFolder.folders {
		fmt.Println(folder.folderName)
	}
}

func newFile(fileName string) *file {
	var fileTest = &file{
		fileName: fileName,
	}
	return fileTest
}

func newFolder(folderName string) *folder {
	var folderTest = &folder{
		folderName: folderName,
		files:      []*file{},
		folders:    make(map[string]*folder),
	}
	return folderTest
}

func (f *folder) addFile(path []string) {
	for i, segment := range path {
		if i == len(path)-1 {
			newFileCreated := newFile(segment)
			f.files = append(f.files, newFileCreated)
		} else {
			fmt.Println("Inside addFile")
			fmt.Println("f:", f.folderName, "segment:", segment)
			f.getFolder(segment).addFile(path[1:])
			return
		}
	}
}

func (f *folder) addFolder(path []string) {
	for i, segment := range path {
		if i == len(path)-1 {
			f.folders[segment] = newFolder(segment)
		} else {
			fmt.Println("Inside addFolder")
			fmt.Println("f:", f.folderName, "segment:", segment)
			f.getFolder(segment).addFolder(path[1:])
		}
	}
}

func (f *folder) getFolder(folderName string) *folder {
	if nextF, ok := f.folders[folderName]; ok {
		return nextF
	} else if f.folderName == folderName {
		fmt.Println("f.folderName:", f.folderName, "folderName:", folderName)
		return f
	} else {
		return &folder{}
	} // cannot happen
}

type file struct {
	fileName string
}

type folder struct {
	folderName string
	files      []*file
	folders    map[string]*folder
}
