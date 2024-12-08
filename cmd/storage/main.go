package main

import (
	"fmt"

	"github.com/Mixai1Ba1/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("hello", st)
}
