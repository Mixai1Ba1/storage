package main

import (
	"fmt"
	"log"

	"github.com/Mixai1Ba1/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("nyka"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("it is restored", restoredFile)
}
