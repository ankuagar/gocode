package main

import (
	"fmt"
	"os"
)

func main() {
	wd, _ := os.Getwd()
	fmt.Println(wd)
	wd, _ = os.Executable()
	fmt.Println(wd)
}
