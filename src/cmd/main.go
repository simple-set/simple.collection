package main

import (
	"fmt"
	"simple.collection/src/version"
)

func main() {
	fmt.Printf("%s/%s\n", version.Name, version.Version)
}
