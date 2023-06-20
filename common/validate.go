package common

import "github.com/klassmann/cpfcnpj"

func ValidatingCPF(cpf string) bool {
	cpf_clear := cpfcnpj.Clean(cpf)
	isValid := cpfcnpj.ValidateCPF(cpf_clear)
	return isValid
}

func ValidatingCNPJ(cnpj string) bool {
	cnpj_clear := cpfcnpj.Clean(cnpj)
	isValid := cpfcnpj.ValidateCNPJ(cnpj_clear)
	return isValid
}

