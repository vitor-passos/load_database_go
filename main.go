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
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	args := os.Args
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
