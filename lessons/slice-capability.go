package main

import "fmt"

func main() {
	// インデックス3つのスライスが導入された
	// [0から数えてn番目か:1から数えてn番目か]
	// [開始:終了:容量]

	items := []string{"A", "B", "C", "D", "E", "F"} // len:6 cap:6
	fmt.Println(len(items), cap(items))
	myItems := items[2:4]
	fmt.Println(myItems) // CD

	// 容量is何？
	// スライスのMaxの長さ

	// スライスには要素を追加できる
	// append()
	items = append(items, "G")
	fmt.Println(items)
	fmt.Println(len(items), cap(items)) // len:6 cap6 -> cap12

	// 容量を超えたら、現在の容量 x 2が再度確保されます！
	// 作成したいスライスの容量が決まっている時
	// 例:野球は9人なので、スライスの容量って9でいい
	// パフォーマンスの為！
}
