package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO customer(id, name) VALUES('john', 'John')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() { // loop query result
		var id, name string
		err := rows.Scan(&id, &name) // match the select field position
		if err != nil {
			panic(err)
		}
		fmt.Println("Customer ID:", id)
		fmt.Println("Customer Name:", name)
	}

	fmt.Println("Success Fetching Customer")
}

func TestQuerySqlComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() { // loop query result
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt) // match the select field position
		if err != nil {
			panic(err)
		}
		fmt.Println("Customer ID:", id)
		fmt.Println("Customer Name:", name)
		if email.Valid {
			fmt.Println("Customer Email:", email)
		}
		fmt.Println("Customer Balance:", balance)
		fmt.Println("Customer Rating:", rating)
		if birthDate.Valid {
			fmt.Println("Customer Birth Date:", birthDate)
		}
		fmt.Println("Customer Married:", married)
		fmt.Println("Customer Created At:", createdAt)
	}

	fmt.Println("Success Fetching Customer")
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #" // sql injection
	password := "wrongpassword"

	query := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' limit 1"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		rows.Scan(&username)
		fmt.Println("Login Success:", username)
	} else {
		fmt.Println("Login Failed")
	}
}

func TestSqlInjectionParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #" // sql injection
	password := "wrongpassword"

	query := "SELECT username FROM user WHERE username = ? AND password = ? limit 1"
	rows, err := db.QueryContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		rows.Scan(&username)
		fmt.Println("Login Success:", username)
	} else {
		fmt.Println("Login Failed")
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "john"
	password := "john"

	query := "INSERT INTO user(username, password) VALUES(?, ?)"
	_, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "john@doe.id"
	comment := "test comment"

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	result, err := db.ExecContext(ctx, query, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Comment with ID:", insertId)
}

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 1; i <= 10; i++ {
		email := "john+" + strconv.Itoa(i) + "@doe.id"
		comment := "comment # " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success insert comment with ID:", lastInsertId)
	}
}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments(email, comment) VALUES(?, ?)"
	for i := 1; i <= 10; i++ {
		email := "john+" + strconv.Itoa(i) + "@doe.id"
		comment := "comment # " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Success insert comment with ID:", lastInsertId)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		panic(err)
	}
}
