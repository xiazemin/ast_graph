package dot

import (
	"fmt"
	"log"
	"os"
)

// WriteToFile  write the svg to file
func WriteToFile(path, name, content string) {
	fname := path + name
	var f *os.File
	var err error
	if pathExists(fname) {
		f, err = os.OpenFile(fname, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	} else {
		f, err = os.Create(fname)
	}
	defer f.Close()

	fmt.Println(fname, err)

	if err != nil {
		log.Println(err.Error())
	}

	_, err = f.Write([]byte(content))
	if err != nil {
		log.Println(err.Error())
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	fmt.Println(err)
	return false
}
