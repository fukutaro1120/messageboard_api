package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"message-board-api/model"
	"message-board-api/repository"

	_ "github.com/go-sql-driver/mysql"
)

// MySQLArticleRepository ...
type MySQLArticleRepository struct {
	db *sql.DB
}

// NewArticleRepository ...
func NewArticleRepository(db *sql.DB) repository.ArticleRepository {
	return &MySQLArticleRepository{db: db}
}

// func NewArticleRepository(db *sql.DB) *MySQLArticleRepository {
// 	return &MySQLArticleRepository{db: db}
// }

// List ...
func (r *MySQLArticleRepository) List(ctx context.Context) ([]*model.Article, error) {
	db, err := sql.Open("mysql", "ito:pass1234@tcp(127.0.0.1:3306)/messageboard")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM articles")
	if err != nil {
		return nil, err
	}
	as := []*model.Article{}
	for rows.Next() {
		var article model.Article
		err = rows.Scan(&article.ID, &article.Name, &article.Message)
		if err != nil {
			return nil, err
		}
		as = append(as, &article)
	}
	return as, nil
}

// Create ...
func (r *MySQLArticleRepository) Create(ctx context.Context, sa *model.Article) error {
	db, err := sql.Open("mysql", "ito:pass1234@tcp(127.0.0.1:3306)/messageboard")
	if err != nil {
		fmt.Printf("create err = %v", err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO articles(Name, Message) VALUES(?,?)")
	if err != nil {
		return err
	}
	fmt.Printf("create err = %v", err)
	_, err = stmt.Exec(sa.Name, sa.Message)
	if err != nil {
		fmt.Printf("create err = %v", err)
	}
	return nil
}

// Delete ...
func (r MySQLArticleRepository) Delete(ctx context.Context, sa *model.Article) error {
	db, err := sql.Open("mysql", "ito:pass1234@tcp(127.0.0.1:3306)/messageboard")
	if err != nil {
		fmt.Printf("delete err = %v", err)
	}
	defer db.Close()
	stmt, err := db.Prepare("DELETE FROM articles WHERE id=?")
	if err != nil {
		fmt.Printf("delete err = %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(sa.ID)
	if err != nil {
		fmt.Printf("delete err = %v", err)
	}
	return nil
}

// Update ...
func (r MySQLArticleRepository) Update(ctx context.Context, sa *model.Article) error {
	db, err := sql.Open("mysql", "ito:pass1234@tcp(127.0.0.1:3306)/messageboard")
	if err != nil {
		fmt.Printf("delete err = %v", err)
	}
	defer db.Close()
	stmt, err := db.Prepare("UPDATE articles SET Message=? WHERE id=?")
	if err != nil {
		fmt.Printf("delete err = %v", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(sa.Message, sa.ID)
	if err != nil {
		fmt.Printf("delete err = %v", err)
	}
	return nil
}
