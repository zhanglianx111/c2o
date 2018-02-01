package models

import (
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
)

func WriteToFile(content []byte) string {
	fn := getFileName()
	err := ioutil.WriteFile(fn, content, 0644)
	if err != nil {
		return ""
	}
	fmt.Printf("file name: %s\n", fn)
	return fn
}

func getFileName() string {
	uuid := uuid.New()
	filename := "/tmp/" + fmt.Sprint(uuid)

	fmt.Printf("file name: %s\n", filename)
	return filename
}
