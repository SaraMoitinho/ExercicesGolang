package main

import (
	"fmt"
	"time"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func main() {
	// language := "Go"
	// fmt.Printf("%s -> %q -> %v\n", language, language, language)

	/*name := "John"
	age := 35

	fmt.Printf("%s\n", name)
	fmt.Printf("%T", age)

	var x, y int
	fmt.Printf("%T -> %d\n", x, y)

	const (
		North = iota
		East
		South
		West
	)

	fmt.Printf("%v %v\n", East, West)

	for sum := 0; index < 10; index++ {
		sum += index
	}

	tum := 1
	for ; bum < 1000 {
		bum += tum
	}

	fmt.Println("sum is equal to ", sum)

	fmt.Println("sum is equal to ", bum)
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}

	integer := 6
	switch integer {
	case 4:
		fmt.Println("integer <= 4")
		fallthrough
	case 5:
		fmt.Println("integer <= 5")
		fallthrough
	case 6:
		fmt.Println("integer <= 6")
		fallthrough
	case 7:
		fmt.Println("integer <= 7")
		fallthrough
	case 8:
		fmt.Println("integer <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}*/
	// retorna o maior valor entre a e b

	x := 3
	y := 4
	z := 5

	max_xy := max(x, y) // chama a função max(x, y)
	max_xz := max(x, z) // chama a função max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // chama a função max aqui

	xPLUSy, xTIMESy := SumAndProduct(x, y)

	fmt.Printf("%d + %d = %d\n", x, y, xPLUSy)
	fmt.Printf("%d * %d = %d\n", x, y, xTIMESy)
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())

}
