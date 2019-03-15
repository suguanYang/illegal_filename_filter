package files

import (
	"easy_tools/lib"
	"io/ioutil"
	"path"
)

func readFileNameFromDir(dirName string, records *[]string) {
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		lib.ErrorHandle(err)
	}
	for i := 0; i < len(fileInfos); i++ {
		fileInfo := fileInfos[i]
		if IsHiddenFile(fileInfo.Name()) {
			continue
		}
		*records = append(*records, path.Join(dirName, fileInfo.Name()))
		if fileInfo.IsDir() {
			readFileNameFromDir(path.Join(dirName, fileInfo.Name()), records)
		}
	}
}

// WalkThroughDirWithHandler just through dir and handle with file name
func WalkThroughDirWithHandler(dirName string, handler FileCb) {
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		lib.ErrorHandle(err)
	}
	for i := 0; i < len(fileInfos); i++ {
		fileInfo := fileInfos[i]
		if IsHiddenFile(fileInfo.Name()) {
			continue
		}
		fileName := fileInfo.Name()
		fileName = handler(path.Join(dirName, fileName))
		if fileInfo.IsDir() {
			WalkThroughDirWithHandler(fileName, handler)
		}
	}
}
