package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

func main() {
	basis()

	// nilマップ
	var ages map[string]int
	fmt.Println(ages == nil) // true
	fmt.Println(len(ages) == 0) // true
	// nilマップへのlen、検索、delete、rageは安全、しかし
	//ages["carol"] = 21 nilマップへの、代入はpanic

	// equal func
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 0}))

	// dedup func
	dedup()

	// charcount
	charcount()

	// graph
	graphFunc()
}

func basis() {
	//ages := map[string]int {
	//	"alice": 31,
	//	"charlie": 34,
	//}
	ages := make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34

	delete(ages, "charlie")// ages["charlie"]の要素を取り除く
	ages["bob"] = ages["bob"] + 1
	ages["bob"] += 1
	ages["bob"]++
	//_ = &ages["undefine"] // コンパイルエラー

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	//var names []string
	names := make([]string, 0, len(ages))// 事前に必要な容量が分かっているので、上より効率的
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

func equal(x, y map[string]int)bool {
	if len(x) != len(y) {
		return false
	}
	for xk, xv := range x {
		if yv, ok := y[xk];!ok || xv != yv {
			return false
		}
	}
	return true
}

func dedup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func charcount() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[r]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func graphFunc() {
	addEdge("from", "to")
	addEdge("first", "second")
	fmt.Println(graph, hasEdge("from", "to"), graph["a"]["b"])
}