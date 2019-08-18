package main

import "fmt"

func main() {
	err := Parse("path/to")
	if err != nil {
		fmt.Println(err)
	}
}

func Parse(input string) (err error) {
	defer func() {
		if p := recover();p != nil {
			err = fmt.Errorf("internal err: %v", p)
		}
	}()
	panic(fmt.Sprintf("parse error: %v", input))
}

func soleTitle() (err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil: // パニックなし
		case bailout{}: // 予期するパニック
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)// 予期しないパニック
		}
	}()
	panic(bailout{})
	return nil
}