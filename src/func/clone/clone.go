package clone

import (
	"fmt"
    "html/template"
    // "io"
//    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
    "strings"
//	"os/exec"
	"../../lib/git"
	"../../lib/appconfig"
	"../../lib/util"
)

const funcname = "clone"

var (
	ip string
	user string
	repository map[string]string
	clone string
	url string
	//fpath string
	//download map[string]string
	//upload string
	//pathchange string
)

type Html struct {
	Ip           string
	User         string
	Repository   map[string]string
	Clone        string
}

func init(){

	// ユーザー設定情報取得
	userConfig, err := appconfig.Parse("./config/user.json")
	if err != nil {
		fmt.Println("error ")
	}

	// ユーザー情報セット
	ip = userConfig.Host + ":"+ userConfig.Port
	url = userConfig.Protocol + "://"+ ip
	user = userConfig.Username

	// repository(navvar)セット
	repository = map[string]string{}
	repository = util.CreateRepository(url)

	// clone(navvar)セット
	clone = url + "/clone"
}

func Index(w http.ResponseWriter, r *http.Request) {

	// view情報
	h := Html{
		User:         user,
		Ip:           ip,
		Repository:   repository,
		Clone:        clone,
	}

	fmt.Println("Regist Index")
	tmpl, _ := template.ParseFiles("./resources/view/clone/index.html")
	tmpl.Execute(w, h)
}

func Regist(w http.ResponseWriter, r *http.Request) {

    if r.Method != "POST" {
        http.Error(w, "Allowed POST method only", http.StatusMethodNotAllowed)
        return
    }

    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }

	r.ParseForm()
	url := r.Form["url"][0]
	account := r.Form["account"][0]
	password := r.Form["password"][0]
	url = strings.Replace(url, "http://", "", 1)

	cloneUrl := "http://" + account + ":" + password + "@" + url

    // あらかじめ戻り先を絶対パスに展開しておく
    prev, err := filepath.Abs(".")
    if err != nil {
        return // ERROR
    }
    defer os.Chdir(prev)

    // ディレクトリ移動
    os.Chdir("repository")

    // clone
    git.Clone(cloneUrl)

	// ディレクトリ移動
    os.Chdir(prev)

	// リダイレクト
	http.Redirect(w, r, "/clone/", http.StatusFound)
}
