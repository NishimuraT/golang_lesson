package main

import (
	"image/color"
	"math"
	"sync"
	"time"
)

// ## 構造体埋め込みによる型の合成
type Point struct {
	X, Y float64
}
func (p *Point)Distance(q *Point) float64 {
	// ColoredPointから呼ばれても、ColoredPointの値は参照できない
	return math.Hypot(q.X- p.X, q.Y- p.Y)
}
type ColoredPoint struct {
	Point
	Color color.RGBA
}
// ## 例：構造体埋め込みによる型の合成
type Cache struct {
	sync.Mutex
	mapping map[string]string
}

// ## nilは正当なレシーバ値
type Values map[string][]string
func (v Values)Get(key string) string {
	if vs := v[key];len(vs) > 0 {
		return vs[0]
	}
	return ""
}
func (v Values)Add(key, value string) {
	v[key] = append(v[key], value) // vがnilの可能性があり、パニックの可能性がある
}

// ## メソッド値
type Rocket struct {}
func (r *Rocket)Launch() {}

func main() {
	// ## 構造体埋め込みによる型の合成
	var a ColoredPoint
	b := ColoredPoint{
		Point: Point{X: 1, Y: 1},
		Color: color.RGBA{
			255, 0, 0, 255,
		},
	}
	a.X = 1 // Xに代入できる
	a.Distance(&b.Point) // Pointのレシーバーが使える

	// ## 例：構造体埋め込みによる型の合成
	cache := Cache{mapping: map[string]string{}}
	cache.Lock()
	cache.mapping["key"] = "value"
	cache.Unlock()

	// ## メソッド値
	p := &Point{X: 1, Y: 1}
	distanceMethod := p.Distance
	distanceMethod(&Point{X: 2, Y: 2})

	r := &Rocket{}
	time.AfterFunc(10 * time.Second, r.Launch)
}