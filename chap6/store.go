package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) ([]Post, error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []Post
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.Id, &post.Content, &post.Author); err != nil {
			return nil, err
		}
		result = append(result, post)
	}
	return result, nil
}

func GetPost(id int) (Post, error) {
	var post Post
	row := Db.QueryRow("select id, content, author from posts where id = $1", id)
	if err := row.Scan(&post.Id, &post.Content, &post.Author); err != nil {
		return Post{}, err
	}
	return post, nil

}

func (post *Post) Create() error {
	q := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
}

func (post *Post) Update() error {
	_, err := Db.Exec("update posts set content = $2, author = $3 where id = $1",
		post.Id, post.Content, post.Author,
	)
	if err != nil {
		return err
	}
	return nil
}

func (post *Post) Delete() error {
	_, err := Db.Exec("delete from posts where id = $1", post.Id)
	if err != nil {
		return err
	}
	return nil
}
