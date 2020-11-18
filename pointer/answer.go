package main

import "fmt"

// メモリは区分けされている
// 部屋がある
// 部屋には大きさがある
// 家具は無限に入れられない
// メモリも物理的な容量は決まっている
// OSがよしなにやっている
// アドレス = 住所
// マンションに例える
// たくさんの部屋がある
// 部屋番号がアドレス
// ポインタを使うと、メモリアドレス上の値を書き換えられる！

func main() {
	var i int
	// Intのポインタ型を宣言
	// int型とは別物。代入不可
	// ポインタ型のnil。int型は0が入る。stringは空文字。
	var ip *int
	fmt.Println(i, ip) // 0 <nil>

	// 注意
	// nilのポインタに操作をするとPanic

	// 変数からポインタを取得する
	// アンパサンドを使う
	ip = &i         // 変数から部屋番号を除く
	fmt.Println(ip) // 0xc0000b2008

	// ポインタから値を取得する
	fmt.Println(*ip) // 0

	// 変数にアスタリスクをつける場合と
	// 型にアスタリスクを付ける場合では挙動に違いあり

	// 型にアスタリスク
	// int型のポインタを宣言
	// var ip2 *int
	// 変数にアスタリスク
	fmt.Println(*ip) // ポインタ変数の値を取得

	var ip3 = new(int)
	fmt.Println(*ip3)

	*ip3 = 12
	fmt.Println(ip3)
	fmt.Println(*ip3)
}
