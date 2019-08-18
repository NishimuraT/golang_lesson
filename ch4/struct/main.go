package main

import (
	"bytes"
	"fmt"
	"image/gif"
	"sync"
	"time"
)

// Name, Addressはまとめる事ができる
type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerId     int
}

type tree struct {
	value int
	left, right *tree //合成型の値は、自分自身を持つ事が出来ませんが、ポインタなら出来ます
}

type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port int
}

//type Circle struct {
//	Center Point
//	Radius int
//}
//type Wheel struct {
//	Circle Circle
//	Spokes int
//}

// 無名フィールドの埋め込み
// 無名フィールのフィールド名は、型名で決まる。名前が衝突するので、二つの無名フィールドを持つ事はできません。
// 可視性も型名で決まります
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

var dilbert Employee

func EmployeeById(id int) *Employee {
	return &dilbert
}

func main() {
	dilbert.Salary -= 5000

	// ポインタを通してフィールドへアクセス
	position := &dilbert.Position
	*position = "Senior" + *position

	// ポインタでもドット表記でアクセスできる
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Position += " (proactive team player)"

	// empty struct（空構造体）フィールドを持たない構造体、boolの代わりに使う事がある
	seen := make(map[string]struct{})
	if _, ok := seen["a"]; !ok {
		seen["a"] = struct{}{}// boolと比べて、メモリを僅かに節約
	}

	// 省略されたフィールドはゼロ値が設定される
	anim := gif.GIF{LoopCount: 3}
	fmt.Println("anim=", anim)

	// 構造体はポインタを通して使用されるのが普通
	pp1 := &Point{1, 2}
	pp2 := new(Point)
	*pp2 = Point{1, 2}
	fmt.Printf("pp1==pp2:%v\tpp1=%v\tpp2=%v\n", pp1 == pp2, pp1, pp2)

	p := Point{1, 2}
	q := Point{2, 1}
	// 構造体は比較可能、次は同じ、フィールドを順番に比較する
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(q == q)

	// 比較可能なので、mapのキー型として使える
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++

	// 構造体のゼロ値は、各フィールドのゼロ値で構成される
	// 例えば、bytes.Bufferの初期値は、空バッファとして使える状態
	var b bytes.Buffer
	fmt.Fprintf(&b, "world!")
	// 例えば、sync.Mutexのゼロ値は、ロックされていないmutexとして使える状態
	var mutex sync.Mutex
	mutex.Unlock()

	// 無名フィールドと埋め込み、構造体リテラル
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	// 構造体リテラルは省略出来ない
	//w = Wheel{8, 8, 5, 20} // コンパイルエラー
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Point: Point{ X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w)
}

// 構造体の引数はポインタでないと、値渡しになる
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}