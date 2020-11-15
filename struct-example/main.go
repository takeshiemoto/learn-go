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
type report struct {
	sol         int
	temperature temperature
	location    location
}

func main() {
	l := location{-4.5567, 137.1111}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: l}
	fmt.Printf("%+v\n", report) // {sol:15 temperature:{high:-1 low:-78} location:{lat:-4.5567 long:137.1111}}
	fmt.Println(report.temperature.high)
	fmt.Println(t.average())
	fmt.Println(report.temperature.average())
}
