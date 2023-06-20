package transform

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"load_database_go/common"
	"load_database_go/models"
	"strconv"
	"strings"
)

const nWorkers = 8

func createStructClientInfos(infos []string) models.ClientInfos {
	cpf := infos[0]
	mostFrequentShop := "null"
	lastBuyShop := "null"
	private := common.StringToBoolean(infos[1])
	incomplete := common.StringToBoolean(infos[2])
	lastBuyDate := strings.ToLower(infos[3])
	averageTicket, err := strconv.ParseFloat(infos[4], 64)
	if err != nil {
		averageTicket = 0.0
	}
	lastBuyTicket, err := strconv.ParseFloat(infos[5], 64)
	if err != nil {
		lastBuyTicket = 0.0
	}
	if common.ValidatingCNPJ(infos[6]) {
		mostFrequentShop = infos[6]
	}
	if common.ValidatingCNPJ(infos[7]) {
		lastBuyShop = infos[7]
	}
	client_info_struct := models.ClientInfos{cpf, private, incomplete, lastBuyDate, averageTicket, lastBuyTicket, mostFrequentShop, lastBuyShop}
	return client_info_struct
}

func ReadFileToDatabase(reader *bufio.Reader, conn *sql.DB) {
	count := 0
	lineInputChan := make(chan []byte, nWorkers)
	for i := 0; i < nWorkers; i++ {
		go worker(lineInputChan, conn)
	}
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println(err)
			return
		}
		if count != 0 {
			lineInputChan <- []byte(line)
		}
		count++
	}
}

func worker(lineInput chan []byte, conn *sql.DB) {
	for line := range lineInput {
		infos := getInfosLine(line)
		client_infos_struct := createStructClientInfos(infos)
		_, err := models.Insert(client_infos_struct, conn)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func getInfosLine(line []byte) []string {
	lineString := string(line)
	var infosClient []string
	stringMod := strings.Replace(lineString, ",", ".", -1)
	splitString := strings.Split(stringMod, " ")
	for _, info := range splitString {
		if strings.TrimSpace(info) != "" {
			infosClient = append(infosClient, info)
		}
	}
	return infosClient
}
