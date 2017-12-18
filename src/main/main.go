package main

import (
	"net/http"
	"../func/clone"
	"../func/files"
	"../func/branchs"
)

func main() {

	// ユーザー設定情報取得
//	userConfig, err := appconfig.Parse("./config/user.json")
//	if err != nil {
//		fmt.Println("error ")
//	}

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))

	http.HandleFunc("/clone/", clone.Index)
	http.HandleFunc("/clone/regist/", clone.Regist)

	http.HandleFunc("/files/", files.Index)
	http.HandleFunc("/branchs/", branchs.Index)

	http.ListenAndServe(":12000", nil)
}
