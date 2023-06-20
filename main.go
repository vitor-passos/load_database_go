package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"load_database_go/common"
	"load_database_go/db"
	"load_database_go/models"
	"io"
	"os"
	"strconv"
	"strings"
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
	getDataBaseClients(reader, conn)

	elapsed := time.Since(startTime)
	fmt.Println("Database Criada com Sucesso em ", elapsed)
}

func createStructClientInfos(infos []string) models.ClientInfos {
	cpf := infos[0]
	loja_mais_frequente := "null"
	loja_ultima_compra := "null"
	private := common.StringToBoolean(infos[1])
	incompleto := common.StringToBoolean(infos[2])
	data_ultima_compra := strings.ToLower(infos[3])
	ticket_medio, err := strconv.ParseFloat(infos[4], 64)
	if err != nil {
		ticket_medio = 0.0
	}
	ticket_ultima_compra, err := strconv.ParseFloat(infos[5], 64)
	if err != nil {
		ticket_ultima_compra = 0.0
	}
	if common.ValidatingCNPJ(infos[6]) {
		loja_mais_frequente = infos[6]
	}
	if common.ValidatingCNPJ(infos[7]) {
		loja_ultima_compra = infos[7]
	}
	client_info_struct := models.ClientInfos{0, cpf, private, incompleto, data_ultima_compra, ticket_medio, ticket_ultima_compra, loja_mais_frequente, loja_ultima_compra}
	return client_info_struct
}

func getDataBaseClients(reader *bufio.Reader, conn *sql.DB) {
	count := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println(err)
			return
		}
		if count != 0 {
			position := getPositionOfWords(line)
			infos := getInfosLine(line, position)
			client_infos_struct := createStructClientInfos(infos)
			_, err := models.Insert(client_infos_struct, conn)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
		count++
	}
}

func getPositionOfWords(line []byte) []int {
	var positions []int
	for i := 1; i < len(line); i++ {
		if line[i] != 32 && line[i-1] == 32 {
			positions = append(positions, i)
		}
	}
	return positions
}

func getInfosLine(line []byte, positions []int) []string {
	initial := 0
	var infos_client []string
	lineString := string(line)
	positions = append(positions, len(lineString))
	fmt.Println(lineString)
	lineMod := strings.Replace(lineString, ",", ".", -1)
	for _, element := range positions {
		word := lineMod[initial:element]
		fmt.Println(word)
		infos_client = append(infos_client, strings.Fields(word)[0])
		initial = element
	}
	return infos_client
}
