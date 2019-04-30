//°C和°F的转换
package main

import "fmt"

func main() {
	const freezeF = 32.0
	const bilingF = 212.0
	fmt.Printf("The %v °F is %v °C\n", freezeF, fToC(freezeF))
	fmt.Printf("The %v °F is %v °C\n", bilingF, fToC(bilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
