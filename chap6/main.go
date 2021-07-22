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
	post := Post{Content: "hello", Author: "nobishii"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, err := GetPost(post.Id)
	if err != nil {
		panic(err)
	}
	fmt.Println(readPost)

	readPost.Content = "HeyHeyHey"
	readPost.Author = "nobiine"

	if err := readPost.Update(); err != nil {
		panic(err)
	}

	posts, err := Posts(10)
	if err != nil {
		panic(err)
	}
	fmt.Println(posts)

	if err := readPost.Delete(); err != nil {
		panic(err)
	}

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
