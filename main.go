package main

import (
	"fmt"

	_ "github.com/HectorHernandezMarques/go_module/test_pkg"
)

func init() {
	fmt.Println("go_module init")
}

func main() {
	fmt.Println("Hello World!")
}
