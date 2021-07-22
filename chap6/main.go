package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	p := Post{Id: 0, Content: "hello", Author: "nobishii"}
	store(p, "test")
	var q Post
	load(&q, "test")
	fmt.Println(q)
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	if err := encoder.Encode(data); err != nil {
		log.Print(err)
		return
	}
	if err := os.WriteFile(filename, buffer.Bytes(), 0644); err != nil {
		log.Print(err)
		return
	}
}

func load(data interface{}, filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Print(err)
		return
	}
	defer f.Close()

	dec := gob.NewDecoder(f)
	if err := dec.Decode(data); err != nil {
		log.Print(err)
		return
	}
}
