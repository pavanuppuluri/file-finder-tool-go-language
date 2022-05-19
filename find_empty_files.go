package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory name to check for empty files")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

  fmt.Println("Below are the empty files");
  
	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			fmt.Println(name)
		}
	}
}
