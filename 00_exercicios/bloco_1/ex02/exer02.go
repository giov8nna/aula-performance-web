// Faça um programa que exiba as tabelas verdade com expressões && e ||

package main

import "fmt"

func main() {
	fmt.Println("|     Tabela Verdade do &&    |")
	fmt.Println("|  Expressão     |    Valor   |")
	fmt.Println("| true && true   |", true && true, "      |")
	fmt.Println("| true && false  |", true && false, "     |")
	fmt.Println("| false && true  |", false && true, "     |")
	fmt.Println("| false && false |", false && false, "     |")

	fmt.Println("\n|     Tabela Verdade do ||    |")
	fmt.Println("|  Expressão     |    Valor   |")
	fmt.Println("| true || true   |    true    |")
	fmt.Println("| true || false  |    true    |")
	fmt.Println("| false || true  |    true    |")
	fmt.Println("| false || false |   false    |")
}
