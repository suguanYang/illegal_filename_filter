package main

import (
	"illegal_filename_filter/files"
)

func main() {
	files.WalkThroughDirWithHandler("./", files.ChangeFileNameByIllegaCharacters)
}
