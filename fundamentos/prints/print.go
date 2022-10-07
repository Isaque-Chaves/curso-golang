package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516

	//fmt.Println("O valor de x é" +  x) // não compila está tentando somar string com num
	xs := fmt.Sprint(x) // convertendo a var x para string
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x, "!!!")

	// Printf = print formatada
	fmt.Printf("O valor de x é %.2f.", x) //%.2f = tipo float com 2 casa decimais

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	// usar cada letra correta pra cada tipo de var
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	// %v serve para vários tipos diferentes de var
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
