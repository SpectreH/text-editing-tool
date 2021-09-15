package errors

import (
	"log"
	"os"
)

var Type1 string = "Error! Missing required text files."
var Type2 string = "Error! No such file or directory."
var Type3 string = "Error! Wrong function syntax."

// prints basic error messages
func PrintError(errType int) {

	if errType == 1 {
		log.Println(Type1)
	} else if errType == 2 {
		log.Println(Type2)
	} else if errType == 3 {
		log.Println(Type3)
	}

	os.Exit(1)
}
