package models

type ClientInfos struct {
	ID                      int64   `json:"id"`
	Cpf                     string  `json:"cpf"`
	Private                 bool    `json:"private"`
	Incompleto              bool    `json:"imcompleto"`
	Data_da_ultima_compra   string  `json:"data_da_ultima_compra"`
	Ticket_medio            float64 `json:"ticket_medio"`
	Ticket_da_ultima_compra float64 `json:"ticket_da_ultima_compra"`
	Loja_mais_frequente     string  `json:"loja_mais_frequente"`
	Loja_da_ultima_compra   string  `json:"loja_da_ultima_compra"`
}

