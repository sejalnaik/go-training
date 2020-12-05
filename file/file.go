package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var isAdded bool = false

func main() {

	startPath := "."
	rootFolder := newFolder(startPath)

	visit := func(path string, info os.FileInfo, err error) error {

		isAdded = false
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

	/*fmt.Println(rootFolder.folderName)
	fmt.Println(rootFolder)
	for _, folder := range rootFolder.folders {
		fmt.Println("folder name:", folder.folderName)
		fmt.Println("folders inside:", folder.folders)
		for _, subFolder := range folder.folders {
			fmt.Println("subfolder name:", subFolder.folderName)
			fmt.Println("subfolders inside:", subFolder.folders)
		}

	}*/

}

func displayFolder(f *folder) {
	if len(f.folders) != 0 {

	}
}

func (f *folder) addFile(path []string) {
	for i, segment := range path {
		if isAdded {
			return
		}
		if f.folderName == segment {
			continue
		}
		if i == len(path)-1 {
			newFileCreated := newFile(segment)
			f.files = append(f.files, newFileCreated)
			isAdded = true
			return
		} else {
			f.getFolder(segment).addFile(path[1:])
			return
		}
	}
}

func (f *folder) addFolder(path []string) {
	for i, segment := range path {
		if isAdded {
			return
		}
		if f.folderName == segment {
			continue
		}
		if i == len(path)-1 {
			f.folders[segment] = newFolder(segment)
			isAdded = true
			return
		} else {
			f.getFolder(segment).addFolder(path[1:])
		}
	}
}

func (f *folder) getFolder(folderName string) *folder {
	return f.folders[folderName]
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

type file struct {
	fileName string
}

type folder struct {
	folderName string
	files      []*file
	folders    map[string]*folder
}
