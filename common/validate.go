package common

import "github.com/klassmann/cpfcnpj"

func ValidatingCPF(cpf string) bool {
	cpfClear := cpfcnpj.Clean(cpf)
	isValid := cpfcnpj.ValidateCPF(cpfClear)
	return isValid
}

func ValidatingCNPJ(cnpj string) bool {
	cnpjClear := cpfcnpj.Clean(cnpj)
	isValid := cpfcnpj.ValidateCNPJ(cnpjClear)
	return isValid
}
