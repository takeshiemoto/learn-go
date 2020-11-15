package main

import "fmt"

type celsius float64

type temperature struct {
	high, low celsius
}

// 型にメソッドを定義
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

type location struct {
	lat, long float64
}

func (l location) days() int {
	return 5
}

type sol int

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

type report struct {
	sol
	//temperature temperature
	temperature
	location location
}

// メソッドの衝突回避
// 上位の型で利用したい方を返す
// 上位の方が優先される
func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

func main() {
	l := location{-4.5567, 137.1111}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: l}
	fmt.Printf("%+v\n", report) // {sol:15 temperature:{high:-1 low:-78} location:{lat:-4.5567 long:137.1111}}
	fmt.Println(report.temperature.high)
	fmt.Println(t.average())
	//fmt.Println(report.temperature.average())
	fmt.Println(report.average()) // temperatureに定義されたメソッドがreportに転送された
	fmt.Println(report.high)      // メソッドだけでなくフィールドも転送される
	report.high = 32
	fmt.Println(report.temperature.high)
	fmt.Println(report.days(1445))
}
