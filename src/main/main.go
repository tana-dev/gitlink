package main

import (
	"net/http"
	"../func/clone"
	"../func/files"
	"../func/branchs"
)

func main() {

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources/"))))

	http.HandleFunc("/clone/", clone.Index)
	http.HandleFunc("/clone/regist/", clone.Regist)

	http.HandleFunc("/files/", files.Index)
	http.HandleFunc("/branchs/", branchs.Index)

	http.ListenAndServe(":12000", nil)
}
