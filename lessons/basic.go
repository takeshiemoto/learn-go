package main

import (
	"fmt"
	"time"
)

// 長さが決まっている配列を定義しなさい
var arrayA = [3]int{1, 2, 3}

// 総素数から推論される配列を定義しなさい
var arrayB = [...]int{1, 2, 3}

// 配列のインデックスと値を指定した配列を定義しなさい
var arrayC = [...]int{1: 1, 3: 3, 9: 9}

// 配列はデータの数まで型情報に含まれるから固定される
// スライスはデータの数を型情報に含まないため増減させることが可能

// スライスを定義せよ
var sliceA = []int{1, 2, 3}

// マップを定義せよ
// 空のmap
var m map[string]int

// 値が入ったmap
var mapA = map[string]int{
	"John": 13,
	"Paul": 21,
}

type Entries = map[string]string

var entries = Entries{
	"2018-10-01": "結婚記念日",
	"2019-05-24": "長男の誕生日",
	"2019-05-29": "友人の誕生日",
	"2022-07-04": "次男の誕生日",
}

const layout = "2006-01-02"

func findEntriesByMonth(entries Entries, year int, month int) ([]string, error) {
	result := make([]string, 0)

	for key, value := range entries {
		t, err := time.Parse(layout, key)
		if err != nil {
			return nil, err
		}
		if t.Year() == year && int(t.Month()) == month {
			result = append(result, value)
		}
	}
	return result, nil
}

func main() {
	results, err := findEntriesByMonth(entries, 2019, 5)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", results)
}
