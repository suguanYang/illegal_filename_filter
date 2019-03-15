package lib

import (
	"log"
)

// ErrorHandle deal with err
func ErrorHandle(err error) {
	log.Fatal(err)
}
