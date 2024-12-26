package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoJoão := ContaCorrente{titular: "João", numeroAgencia: 589, numeroConta: 12345, saldo: 125.5}
	fmt.Println(contaDoJoão)

	contaDoVinicius := ContaCorrente{"Vinícius", 590, 123456, 200}
	fmt.Println(contaDoVinicius)
}
