package regist

import (
	"fmt"
    "html/template"
    // "io"
    // "io/ioutil"
    // "net/http"
    "os"
    "path/filepath"
    "strings"
	// "os/exec"
	// "../git"
)

type Html struct {
	FileinfoList [][]string
	Breadcrumbs  map[string]string
	User         string
	Ip           string
}

func Index(w http.ResponseWriter, r *http.Request) {

//	var url string
//	var currentDir string

	var ip string
	var user string
	var url string
	var fileinfoList [][]string
	var breadcrumbs map[string]string
	var fpath string
	var fname string
	var download map[string]string
	var upload string
	var pathchange string

	// ユーザー設定情報取得
	userConfig, err := appconfig.Parse("./config/user.json")
	if err != nil {
		fmt.Println("error ")
	}

	// ユーザー情報セット
	ip = userConfig.Host + ":"+ userConfig.Port
	url = userConfig.Protocol + "://"+ ip
	user = userConfig.Username

	// fpath = r.URL.Path
	// fpath1 := r.URL.Path
	// fpath1 = strings.TrimRight(fpath1, "/")

	// pathを取るにはr.URL.Pathで受け取文末のスラッシュを削除
	// fpath = strings.TrimRight(fpath, "/") // 2. Linux
	// fname = filepath.Base(fpath)

	// currentDir, _ = filepath.Abs(".")
	// createDir(currentDir + "/repositories")
    //
	// cloneUrl := "https://github.com/tana-dev/practice.git"
    // os.Chdir(currentDir + "/repositories")
    //
    // git.Clone(cloneUrl)
    // os.Chdir(currentDir)

	// out, _ := exec.Command("git", "branch", "-r").Output()
	// fmt.Println(string(out))

	// tmpl := template.Must(template.ParseFiles("./view/index.html"))
	// tmpl.Execute(w, h)

	fmt.Println("Regist")
	tmpl, _ := template.ParseFiles("./resources/view/regist/index.html")
	tmpl.Execute(w, h)
}
