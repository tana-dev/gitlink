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
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	// var url string
	// var currentDir string

	// url = "http://10.27.145.100:8080/"
	// url = "http://192.168.33.22:8080/"

	// fpath = r.URL.Path
	// fpath1 := r.URL.Path
	// fpath1 = strings.TrimRight(fpath1, "/")

	// pathを取るにはr.URL.Pathで受け取文末のスラッシュを削除
	// fpath = strings.TrimRight(fpath, "/") // 2. Linux
	// fname = filepath.Base(fpath)

    a := "Regist"
    action.a()
	fmt.Println("tanaka")

	// out, _ := exec.Command("git", "branch", "-r").Output()
	// fmt.Println(string(out))

	// tmpl := template.Must(template.ParseFiles("./view/index.html"))
	// tmpl.Execute(w, h)

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
