package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64) inferido pelo compilador

	// forma reduzida de criar uma var, declara var e atribui um valor [recomendado]
	area := PI * m.Pow(raio, 2)

	// uma var sempre tem que ser chamada
	fmt.Println("A área da circunferência é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	// varias var em uma mesma linha de código (var e bool true | var f bool false)
	var e, f bool = true, false
	fmt.Println(e, f)

	// sintaxe reduzida, o compilador não varia durante o programa, ou seja a var 'g' sempre será inteira etc
	g, h, i := 2, false, "epa!"
	fmt.Println(g, h, i)
}
