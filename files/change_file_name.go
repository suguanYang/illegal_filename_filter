package files

import (
	"illegal_filename_filter/lib"
	"os"
	"regexp"
)

type changeFilesNameOptions struct {
	pattern     string
	replacement string
}

func deleteStringAtIndex(str string, index int) string {
	if index == len(str)-1 {
		return str[:index]
	}
	return str[:index] + str[index+1:]
}

func replaceFileNameInPath(path string, options changeFilesNameOptions) string {
	re := regexp.MustCompile(options.pattern)
	illegaIndex := 0
	baseIndex := 0
	converted := path
	for i := 0; i < len(path); i++ {
		here := string(path[i])
		if re.MatchString(here) {
			if i-illegaIndex == 1 && i > 1 {
				converted = deleteStringAtIndex(converted, i-baseIndex)
				baseIndex++
			}
			illegaIndex = i
		}
	}

	return re.ReplaceAllString(converted, options.replacement)
}

// ChangeFileNameByPattern change the gives file name
func ChangeFileNameByPattern(filePath string, options changeFilesNameOptions) string {
	if IsFileExists(filePath) {
		newPath := replaceFileNameInPath(filePath, options)
		if newPath != filePath {
			if err := os.Rename(filePath, newPath); err != nil {
				lib.ErrorHandle(err)
			}
			return newPath
		}
	}
	return filePath
}

// ChangeFileNameByIllegaCharacters replace the illega characters with _
func ChangeFileNameByIllegaCharacters(filePath string) string {
	options := changeFilesNameOptions{pattern: `[\&\\{\}\\\|\[\]\~\^]`, replacement: "_"}
	return ChangeFileNameByPattern(filePath, options)
}
