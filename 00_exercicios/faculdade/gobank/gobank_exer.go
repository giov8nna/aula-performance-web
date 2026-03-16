package main

import "fmt"

func depositar(saldoConta *float64, valorDeposito float64) float64 {
	*saldoConta = *saldoConta + valorDeposito
	return *saldoConta
}

func sacar(saldoConta *float64, valorSaque float64) float64 {
	if *saldoConta < valorSaque {
		fmt.Println("Saldo da conta insuficiente: ", *saldoConta)
	} else {
		*saldoConta = *saldoConta - valorSaque
	}
	return *saldoConta
}

func main() {
	var saldoConta float64 = 100.00

	fmt.Println("Você possui: ", saldoConta)

	depositar(&saldoConta, 50.00)

	fmt.Println("Saldo da conta após depósito: ", saldoConta)

	sacar(&saldoConta, 30.00)

	fmt.Println("Saldo da conta após o saque: ", saldoConta)

	sacar(&saldoConta, 300.00)
}
