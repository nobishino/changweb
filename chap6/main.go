package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	data := []byte("Hello, World!\n")

	if err := os.WriteFile("data1", data, 0644); err != nil {
		log.Fatal(err)
	}

	read1, err := os.ReadFile("./data1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(read1))

	f1, err := os.Create("data2")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()
	if _, err := f1.Write(data); err != nil {
		log.Fatal(err)
	}
	f1.Sync()

	f2, err := os.Open("data2")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	io.Copy(os.Stdout, f2)

}
