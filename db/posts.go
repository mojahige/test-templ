package db

type Post struct {
	ID        int
	Title     string
	Body      string
	CreatedAt string
	UpdatedAt *string
}

func GetAllPosts() ([]Post, error) {
	rows, err := DB.Query("SELECT id, title, body, created_at, updated_at FROM posts")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []Post{}

	for rows.Next() {
		var post Post

		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func GetPost(id int) (Post, error) {
	var post Post

	err := DB.QueryRow("SELECT id, title, body, created_at, updated_at FROM posts WHERE id = ?", id).Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt, &post.UpdatedAt)

	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func CreatePost(title, body string) (int64, error) {
	result, err := DB.Exec("INSERT INTO posts (title, body) VALUES (?, ?)", title, body)

	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func UpdatePost(id int, title, body string) error {
	_, err := DB.Exec("UPDATE posts SET title = ?, body = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?", title, body, id)

	return err
}

func DeletePost(id int) error {
	_, err := DB.Exec("DELETE FROM posts WHERE id = ?", id)

	return err
}

func InitPostsTable() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			body TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
			updated_at TIMESTAMP
		)
	`)

	return err
}
