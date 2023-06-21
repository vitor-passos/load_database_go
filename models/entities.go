package models

type ClientInfos struct {
	Cpf              string  `json:"cpf"`
	Private          bool    `json:"private"`
	Incomplete       bool    `json:"imcompleto"`
	LastBuyDate      string  `json:"data_da_ultima_compra"`
	AverageTicket    float64 `json:"ticket_medio"`
	LastBuyTicket    float64 `json:"ticket_da_ultima_compra"`
	MostFrequentShop string  `json:"loja_mais_frequente"`
	LastBuyShop      string  `json:"loja_da_ultima_compra"`
}
