package main

import (
	"os"
	"fmt"
//	"encoding/json"
)

func usage() {
	fmt.Println("raygo [JSON_SCENE]...")
}

func main() {
	if len(os.Args) > 1 {
		fmt.Println("Hello World")
	} else {
		usage()
		os.Exit(0)
	}
}
