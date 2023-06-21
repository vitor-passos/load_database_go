package transform

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"load_database_go/common"
	"load_database_go/db"
	"load_database_go/models"
	"strings"
)

const nWorkers = 8

func createStructClientInfos(infos []string) (models.ClientInfos, bool) {
	if !common.ValidatingCPF(infos[0]) {
		return models.ClientInfos{}, false
	}
	cpf := infos[0]
	mostFrequentShop := "null"
	lastBuyShop := "null"

	if common.ValidatingCNPJ(infos[6]) {
		mostFrequentShop = infos[6]
	}
	if common.ValidatingCNPJ(infos[7]) {
		lastBuyShop = infos[7]
	}

	return models.ClientInfos{
		Cpf:              cpf,
		Private:          common.StringToBoolean(infos[1]),
		Incomplete:       common.StringToBoolean(infos[2]),
		LastBuyDate:      strings.ToLower(infos[3]),
		AverageTicket:    common.ParseToFloatWithDefault(infos[4], 0.0),
		LastBuyTicket:    common.ParseToFloatWithDefault(infos[5], 0.0),
		MostFrequentShop: mostFrequentShop,
		LastBuyShop:      lastBuyShop,
	}, true
}

func ReadFileToDatabase(reader *bufio.Reader, conn *sql.DB) {
	count := 0
	lineInputChan := make(chan string, nWorkers)

	for i := 0; i < nWorkers; i++ {
		go worker(lineInputChan, conn)
	}

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if count != 0 {
			lineInputChan <- line
		}
		count++
	}
}

func worker(lineInput chan string, conn *sql.DB) {
	for line := range lineInput {
		infos := getInfosLine(line)
		clientInfosStruct, isValid := createStructClientInfos(infos)
		if !isValid {
			continue
		}
		_, err := db.Insert(clientInfosStruct, conn)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func getInfosLine(line string) []string {
	var infosClient []string
	stringMod := strings.Replace(line, ",", ".", -1)
	splitString := strings.Split(stringMod, " ")
	for _, info := range splitString {
		if strings.TrimSpace(info) != "" {
			infosClient = append(infosClient, info)
		}
	}
	return infosClient
}
