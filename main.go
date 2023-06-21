package main

import (
	"bufio"
	"fmt"
	"load_database_go/db"
	"load_database_go/transform"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Println("Starting load_database_go application...")
	conn, err := db.OpenConnection()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	args := os.Args
	if args[1] == "" {
		fmt.Println("Path with file needed.")
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(file)
	fmt.Println("Processing File ", args[1], "this may take a few seconds")
	transform.ReadFileToDatabase(reader, conn)

	elapsed := time.Since(startTime)
	fmt.Println("Successufully entered data within ", elapsed)
}
