package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var idUsuario int = rand.Intn(10000)
	var status bool = false

	if status {
		fmt.Printf("Iniciando Sessão para Usuário %d\n", idUsuario)
	} else {
		fmt.Printf("Falha ao iniciar sessão para Usuário %d\n", idUsuario)
	}

}
