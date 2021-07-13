package main

import "fmt"

func main() {
	/* local variable definition */
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a: %d\n", a)
	fmt.Printf("Before swap, value of b: %d\n", b)

	/* Calling a function to swap the values.
	 * &a indicates pointer to a ie, address of variable a and
	 * &b indicates pointer to b ie. address of variable b.
	 */
	swap(&a, &b)

	fmt.Printf("After swap, value of a: %d\n", a)
	fmt.Printf("After swap, value of b: %d\n", b)
}

func swap(a *int, b *int) {
	var temp int
	temp = *a /* save the value at address x */
	*a = *b   /* put y to x */
	*b = temp /* put temp into y */
}
