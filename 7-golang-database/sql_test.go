package golang_database

import (
	"context"
	"database/sql"
	"fmt"
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
