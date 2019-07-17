package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello, 世界!")
	fmt.Println("GOOS:", runtime.GOOS)
	fmt.Println("GOARCH", runtime.GOARCH)
}
