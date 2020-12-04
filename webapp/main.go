package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	// 静的ファイルの転送処理
	files := http.FileServer(http.Dir("public"))
	// staticで始まるすべてのリクエストをpublicを起点として
	// 残った文字列を名前荷物ファイルを探す
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	// indexは別ファイルだが同一パッケージ内ならばimport不要で参照可能
	mux.HandleFunc("/", index)
	/// error
	mux.HandleFunc("/", err)

	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/session", session)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
