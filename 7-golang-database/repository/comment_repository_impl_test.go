package repository

import (
	golang_database "7-golang-database"
	"7-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "john@repository.id",
		Comment: "Test Comment",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 22)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(comments)
}
