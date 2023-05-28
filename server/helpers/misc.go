package helpers

import "log"

func CheckErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
