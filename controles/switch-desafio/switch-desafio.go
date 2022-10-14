package main

import "fmt"

func refatorarSwitch1(n float64) string {
	nota := int(n)

	switch {
	case nota >= 9:
		return ("A")
	case nota == 8 || nota == 7:
		return ("B")
	case nota == 6 || nota == 5:
		return ("C")
	case nota == 4 || nota == 3:
		return ("D")
	case nota == 2 || nota == 1 || nota == 0:
		return ("E")
	default:
		return ("Nota inválida")
	}
}

func main() {
	fmt.Println(refatorarSwitch1(-1))
}
