package main

import "fmt"

func main() {
	answer := 42
	fmt.Println(&answer)

	address := &answer
	// アドレスから値を参照する
	fmt.Println(*address)

	fmt.Println(*&answer) //42

	fmt.Printf("addressの型は %T です\n", address) //addressの型は *int です

	var manager *string
	oh := "王貞治"

	manager = &oh

	fmt.Println(*manager) // 王貞治

	akiyama := "秋山幸二"
	manager = &akiyama

	fmt.Println(*manager) // 秋山幸二

	// akiyamaをkudoに変更すると...?
	akiyama = "工藤公康"

	fmt.Println(*manager) // 工藤公康

	// managerをデリファレンスして直接書き換えると...?
	*manager = "城島健司"
	fmt.Println(akiyama) // 城島健司!?

	mrHawks := manager
	*mrHawks = "小久保裕紀"
	fmt.Println(akiyama) // 小久保裕紀???

	// mrHawksはmanagerと同じメモリアドレス
	fmt.Println(manager == mrHawks) // true

	saito := "斉藤和巳"
	manager = &saito                // managerにsaitoの参照を入れる
	fmt.Println(manager == mrHawks) // false
}
