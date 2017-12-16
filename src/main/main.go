package main

import (
	"fmt"
	"path/filepath"
	"net/http"
	"os"
	// "os/exec"
	// "./git"
	"../action"
)

func main() {

	// fpath := r.URL.Path
	// fpath = strings.TrimRight(fpath1, "/")
    u, err := url.Parse("http://user1@bing.com/search?q=foo%2fbar&q2=hoge#fragment")
    if err != nil {
      log.Fatal(err)
    }
    path = u.Path
    fmt.Printf("Path: %s\n", u.Path)
    fmt.Printf("RawQuery: %s\n", u.RawQuery)

	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(assetFS()))))

    http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
        if r.URL.Path == "/regist/" {
            action.Regist(w)
        } else if r.URL.Path == "/files/" {
            action.Diff(w)
        } else if r.URL.Path == "/diff/" {
            action.Diff(w)
        }
    })

	http.ListenAndServe(":12000", nil)
}

func createDir(dir string) {

	_, err := os.Stat(dir)
	if err == nil {
		return
	}

	if err := os.MkdirAll(dir, 0777); err != nil {
		fmt.Println(err)
	}
}

func cd(dir string) {

    // あらかじめ戻り先を絶対パスに展開しておく
    prev, err := filepath.Abs(".")
    if err != nil {
        return // ERROR
    }
    defer os.Chdir(prev)

    // ディレクトリ移動
    os.Chdir(dir)
}
