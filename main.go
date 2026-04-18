package main

import (
	"AutoSyncGO/code/db"
	"fmt"
)

func main() {
	fmt.Println("I should be studying ML right now")
	db.Connect("user:password@tcp(127.0.0.1:3306)/dbname")
}
