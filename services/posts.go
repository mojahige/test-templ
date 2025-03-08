package services

import (
	"github.com/mojahige/test-templ/db"
)

func GetAllPosts() ([]db.Post, error) {
	return db.GetAllPosts()
}

func GetPost(id int) (db.Post, error) {
	return db.GetPost(id)
}

func CreatePost(title, body string) (int64, error) {
	return db.CreatePost(title, body)
}

func UpdatePost(id int, title, body string) error {
	return db.UpdatePost(id, title, body)
}

func DeletePost(id int) error {
	return db.DeletePost(id)
}
