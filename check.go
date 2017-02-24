package goutils

import (
	"log"
)

func PanicIf(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func WarnIf(err error) {
	if err != nil {
		log.Println(err)
	}
}
