package models

import (
	"database/sql"
)

func Insert(clientInfos ClientInfos, conn *sql.DB) (id int64, err error) {
	sql := `INSERT INTO client_infos (cpf, private, imcompleto, data_da_ultima_compra, ticket_medio, ticket_da_ultima_compra, loja_mais_frequente, loja_da_ultima_compra) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err = conn.QueryRow(sql, clientInfos.Cpf, clientInfos.Private, clientInfos.Incomplete, clientInfos.LastBuyDate, clientInfos.AverageTicket, clientInfos.LastBuyTicket, clientInfos.MostFrequentShop, clientInfos.LastBuyShop).Scan(&id)
	return
}
