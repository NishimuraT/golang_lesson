package main

import "fmt"

func main() {
	m := [...]string{1: "January", 12: "December"}
	fmt.Println(len(m)) // 13

	months := [...]string{
		1: "January", "February", "March", "April", "May", "June", "July", "August", "September", "Octobar", "November", "December",
	}
	fmt.Println(months)
	Q2 := months[4:7] //[i:j] iからj-1までの３つ
	summer := months[6:9]
	fmt.Println(Q2, len(Q2), cap(Q2))
	fmt.Println(summer, len(summer), cap(summer))
	fmt.Println(months[:]) // 配列全体

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	//fmt.Println(summer[:20]) panic: 範囲外
	fmt.Println(summer[:5])// スライスを拡張（容量の範囲内で）

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // [5, 4, 3, 2, 1, 0]

	s := []int{0, 1, 2, 3, 4, 5}// 暗黙に基底配列を生成し、その配列を指すスライスを生成
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)// [2, 3, 4, 5, 0, 1]

	// nilスライス（nilスライスは基底配列を持たない、長さは０、容量は０）
	var s2 []int // len=0, s==nil
	s2 = nil // len=0, s==nil
	s2 = []int(nil) // len=0, s==nil
	s2 = []int{} // len=0, s!=nil

	// スライスが空であるかを検査する必要がある時
	if len(s2) == 0 {
		fmt.Println("slice is empty")
	}

	// make([]T, len, cap)
	fmt.Println(make([]int, 2)) // 容量が省略された場合は、capとlenが同じになる
	fmt.Println(make([]int, 2, 3))


	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	// todo
	//var x, y []int
	//for i := 0;i < 10;i++ {
	//	y = appendInt(x, i)
	//	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	//	x = y
	//}
	data := []string{"one", "", "three"}
	fmt.Println(nonempty(data))
	fmt.Println(data)

	// stack
	stack := []string{}
	stack = append(stack, "a")// push
	//top := stack[len(stack)-1] //top
	stack = stack[:len(stack)-1]// pop

	// remove
	s1 := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s1, 1))
}

// reverse(nil)は安全
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j;i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// スライスは配列と違い、比較可能ではない（[]byteは、bytes.Equal関数がある）
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// 拡大する余地がある。スライスを拡張する。
		z = x
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)// cap(dist, source)
	}
	z[len(x)] = y
	return z
}

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings
}

func nonempty2(strings []string) []string {
	out := []string{}
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[:i], slice[:i+1])
	return slice[:len(slice)-1]
}

// 順序を維持しなくても良い場合
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice) - 1]
	return slice[:len(slice) - 1]
}