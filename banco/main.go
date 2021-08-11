package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto( conta verificarConta, valorDoBoleto float64 ) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	 Sacar(valor float64) string
}


func main() {
	contaDoDaniel := contas.ContaPoupanca{}
	contaDoDaniel.Depositar(500)
	PagarBoleto(&contaDoDaniel, 100)

	fmt.Println(contaDoDaniel.ObterSaldo())

	contaDaDanielle := contas.ContaCorrente{}
	contaDaDanielle.Depositar(900)
	PagarBoleto(&contaDaDanielle, 100)

	fmt.Println(contaDaDanielle.ObterSaldo())
}