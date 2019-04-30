//变量的生命周期
package main

var global *int

func f() {
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}

func main() {
	
}
