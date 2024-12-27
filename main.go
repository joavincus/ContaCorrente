package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso! "
	} else {
		return "Saldo insuficiente."
	}
}

func main() {
	contaDaSilva := ContaCorrente{}
	contaDaSilva.titular = "Silva"
	contaDaSilva.saldo = 500

	fmt.Println(contaDaSilva.saldo)
	fmt.Println(contaDaSilva.Sacar(600))
	fmt.Println(contaDaSilva.saldo)
}
