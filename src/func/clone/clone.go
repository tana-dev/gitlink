package clone

import (
	"fmt"
    "html/template"
    // "io"
    "io/ioutil"
    "net/http"
    "os"
    "path/filepath"
    "strings"
//	"os/exec"
	"../../lib/git"
	"../../lib/appconfig"
)

type Html struct {
	User         string
	Ip           string
	Repository   map[string]string
	Clone        string
}

func Index(w http.ResponseWriter, r *http.Request) {

	var ip string
	var user string
	var url string
	var fpath string
//	var download map[string]string
//	var upload string
//	var pathchange string
	var repository map[string]string

	// ユーザー設定情報取得
	userConfig, err := appconfig.Parse("./config/user.json")
	if err != nil {
		fmt.Println("error ")
	}

	// ユーザー情報セット
	ip = userConfig.Host + ":"+ userConfig.Port
	url = userConfig.Protocol + "://"+ ip
	user = userConfig.Username

	// レポジトリー取得
	fpath = r.URL.Path
	fpath = strings.Replace(fpath, "/files/", "", 1)
//	fpath = strings.TrimRight(fpath, "/")

	// repositoryセット
	repository = map[string]string{}
	repos := dirwalk("repository")
	for _, rp := range repos {
		var name string
		link := rp
		link = strings.Replace(link, `\`, "/", -1)                // 1.Windows
//		link = url + "/files" + strings.Replace(link, "/", "", 2) // 1.Windows
		link = url + "/files" + strings.Replace(link, "repository", "", 1) + "/" // 2.Linux
		name = filepath.Base(rp)
		repository[link] = name
	}

	h := Html{
		User:         user,
		Ip:           ip,
		Repository:   repository,
	}

	fmt.Println("Regist Index")
	tmpl, _ := template.ParseFiles("./resources/view/clone/index.html")
	tmpl.Execute(w, h)
}

func Regist(w http.ResponseWriter, r *http.Request) {

//	var ip string
//	var fpath string
//	var download map[string]string
//	var upload string
//	var pathchange string

	// ユーザー設定情報取得
//	userConfig, err := appconfig.Parse("./config/user.json")
//	if err != nil {
//		fmt.Println("error ")
//	}

	// ユーザー情報セット
//	ip = userConfig.Host + ":"+ userConfig.Port

	// レポジトリー取得
//	fpath = r.URL.Path
//	fpath = strings.Replace(fpath, "/files/", "", 1)
//	slashNum := strings.Index(fpath, "/")
//	current_repo := fpath[0:slashNum]

    // あらかじめ戻り先を絶対パスに展開しておく
    prev, err := filepath.Abs(".")
    if err != nil {
        return // ERROR
    }
    defer os.Chdir(prev)

    // ディレクトリ移動
    os.Chdir("repository")

	cloneUrl := "http://tanaka-shu:st1127pass@glab.gnavi.co.jp/attribute/attribute.git"
    git.Clone(cloneUrl)

	// ディレクトリ移動
    os.Chdir(prev)

	// リダイレクト
	http.Redirect(w, r, "/clone/", http.StatusFound)
}

func dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	var dpaths []string
	var fpaths []string
	for _, file := range files {
		if 0 != strings.Index(file.Name(), ".") && 0 != strings.Index(file.Name(), "~$") && 0 != strings.Index(file.Name(), "Thumbs.db") {

			f := filepath.Join(dir, file.Name())

			// ファイル存在チェック
			fi, _ := os.Stat(f)
			if fi.IsDir() {
				dpaths = append(dpaths, filepath.Join(dir, file.Name()))
			} else {
				fpaths = append(fpaths, filepath.Join(dir, file.Name()))
			}
		}
	}

	if nil == dpaths && nil != fpaths {
		paths = fpaths
	} else if nil != dpaths && nil == fpaths {
		paths = dpaths
	} else {
		paths = append(dpaths, fpaths...)
	}

	return paths
}
