package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	port := flag.String("port", "5437", "Put your database port")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")

	flag.Parse()

	fmt.Println("host:", *host)
	fmt.Println("port:", *port)
	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
}
