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
	file, err := os.Open("base_teste.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reader := bufio.NewReader(file)
	transform.ReadFileToDatabase(reader, conn)

	elapsed := time.Since(startTime)
	fmt.Println("Database Criada com Sucesso em ", elapsed)
}
