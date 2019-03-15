package main

import (
	"easy_tools/files"
)

func main() {
	files.WalkThroughDirWithHandler("./", files.ChangeFileNameByIllegaCharacters)
}
