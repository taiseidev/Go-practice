package main

import "fmt"

func foo() {
	xi := 1
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T", xf32)
	xs = "書き換えました"
	fmt.Println(xi, xf32, xs, xt, xf)
}

func main() {
	foo()

	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)
	fmt.Println(i, f64, s, t, f)

}
