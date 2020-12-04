package main

import (
	"fmt"
	"net/http"
)

// ユーザーを認証し、クライアントにクッキーを返すハンドラ
func authenticate(w http.ResponseWriter, r *http.Request) {
	// TODO
	// ユーザーを認証する
	// パスワードが一致しているかをチェックする
	// 認証されたら構造体Uを返す
	// クライアントのブラウザにクッキーを埋める
	fmt.Fprintf(w, "Authenticate!!, %s", r.URL.Path[1:])
}
