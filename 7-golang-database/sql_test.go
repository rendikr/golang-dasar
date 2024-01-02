package golang_database

import (
	"context"
	"fmt"
	"testing"
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
