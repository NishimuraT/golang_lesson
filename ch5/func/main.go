package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
)

// デフォルト引数は存在しない、名前で引数を指定する方法はない
// したがって、関数のパラメーターと結果の名前は、呼び出しもとにとって関心がない（doc以外）
// 値渡し、ポインタ、スライス、マップ、関数、チャネルに対する修正は呼び出しもとは修正の影響を受けるかもしれません
func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x -y; return
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}

// 本体がない関数宣言は、Go以外で実装されている関数を示します
//package math
//func Sin(x float64) float64 // アセンブリ言語

func main() {
	// basic
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)

	// 再帰、golang.org/x/net/html
	doc := fetch()
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	// 関数値
	f := add
	fmt.Println(f(2, 3))
	fmt.Printf("%T\n", f)

	// 関数型のnil値
	var fun func(int)int
	if fun != nil {
		fun(3)
	}
	//fun(3) // panic
	// 関数型は比較可能ではなく、mapのキーとして使用できない

	// 関数値を使い、関数の振る舞いを定義する事ができる（as strings.Map(mapping func(rune) rune, s string)）
	rot13 := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+13)%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	fmt.Println(strings.Map(rot13, "'Twas brillig and the slithy gopher..."))

	// 関数リテラル
	s1 := squares()// 関数値は、状態（ローカル変数）をモテる
	s2 := squares()
	fmt.Println(s1())// 1
	fmt.Println(s2())// 1
	fmt.Println(s1())// 4
	fmt.Println(s2())// 4
	fmt.Println(s1())// 9
	fmt.Println(s2())// 9
	fmt.Println(s1())// 16
	fmt.Println(s2())// 16

	// 可変長
	fmt.Println(sum(1, 2, 3, 4))
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	// 遅延関数呼び出し
	defer func() {
		fmt.Println("first define")
	}()
	defer func() {
		fmt.Println("second define")
	}()
}

func fetch() *html.Node {
	resp, err := http.Get("http://gopl.io")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "html parse: %v\n", err)
		os.Exit(1)
	}
	return doc
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// 多値の返り値は、次の様に関数を書いても良い
func multiReturnLog(arg1 string) (string, error) {
	log.Printf("multiReturn %s\n", arg1)
	return multiReturn(arg1)
}

func multiReturn(arg1 string) (string, error) {
	return arg1, nil
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		resp.Body.Close()
		return
	}
	_, err = html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = 0, 0
	return
}

func squares() func() int {
	var x int // 変数の生存期間が、スコープで決まるわけではない
	// 無名関数
	return func() int {
		x++ // クロージャにより、xにアクセスできる
		return x * x
	}
}

func topoSort(m map[string][]string)[]string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string) // 宣言（宣言と代入がまとめられていると、再帰的な実行時にコンパイルエラー）

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

// ループ変数の捕捉
func loopVal() {
	var rmdirs []func()
	//for _, dir := range []string{} { // ループ変数は、各関数値に共有される
	//	os.MkdirAll(dir, 0755)
	//	rmdirs = append(rmdirs, func() {
	//		os.RemoveAll(dir) // 意図通り働かない
	//	})
	//}
	for _, dir := range []string{} {
		dir := dir// 内側のdirを宣言し、外側のdirで初期化する
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
	// ループ変数は、go文、deferでも問題になるかもしれません
}

// 可変長関数
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
// 上と下では明確に型が違う
//func sum(vals []int) int {}

// 遅延関数の呼び出し
func triple(x int) (result int) {
	defer func() { result += x }()
	return x * 2
}
func fileOpen(filenames []string) error {
	//for _, filename := range filenames {
	//	f, err := os.Open(filename)
	//	if err != nil {
	//		return err
	//	}
	//	defer f.Close() // 関数の処理が終わるまで実行されないので、ファイル記述子が枯渇する可能性がある
	//}
	for _, filename := range filenames {
		err := doFile(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
// １つの解決策は別の関数に移す事
func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
func fetch2(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// f.Close()をdeferに設定するのは良くない。エラー処理を行うべきである
	if closeErr := f.Close(); err != nil {
		err = closeErr
	}
	return local, n, nil
}