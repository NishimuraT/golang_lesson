package main

import "fmt"

// 関数のローカル変数のアドレスを返す事は安全
func f() *int {
	v := 1
	return &v
}

func main() {
	// ## basis
	x := 1
	p := &x
	fmt.Println(*p)// 1
	*p = 2
	fmt.Println(x)// 2

	// ## ゼロ値はnil、どの型に対するポインタでも、ゼロ値はnilである

	// ## 構造体のフィールド、配列の要素も変数であり、アドレス化可能

	var a, b int
	fmt.Println(&a == &a, &a == &b, &a == nil)// true, false, false

	// ## 関数のローカル変数
	fmt.Println(f() == f())// false

	// ## new関数
	// new(t)はTのゼロ値へ初期化し、アドレスを返す
	c := new(int)
	*c=2
	fmt.Println(*c)
}
