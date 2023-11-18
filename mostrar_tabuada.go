package main

import "fmt"

func main() {
	numeros := [10]int{1,2,3,4,5,6,7,8,9,10}
	//fmt.Print("Digite um número para ver a tabuada: ")
	//var numero int
	//fmt.Scanln(&numero)

	for a := 0; a < len(numeros); a++ {
		fmt.Println("-----------------")
		elemento := numeros[a]
		for i := 1; i <= 10; i++ {
			resultado := elemento * i
			fmt.Printf("%d x %d = %d\n", elemento, i, resultado)
		}
	}
}
