package branchs

import (
	"fmt"
	"html/template"
//	"io"
//	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"../../lib/git"
	"../../lib/appconfig"
	"../../lib/util"
)

const funcname = "branchs"

var (
	ip string
	user string
	branchList [][]string
	clone string
	sidelink map[string]string
	repository map[string]string
	url string
	fpath string
)

type Html struct {
	Ip           string
	User         string
	Repository map[string]string
	Clone        string
	Sidelink     map[string]string
	BranchList [][]string
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

	// パス取得
	fpath = strings.Replace(r.URL.Path, "/" + funcname + "/", "", 1)
	slashNum := strings.Index(fpath, "/")
	current_repo := fpath[0:slashNum]

	// sidelinkセット
	sidelink = map[string]string{}
	sidelink = util.CreateSidevar(url,current_repo)

    // あらかじめ戻り先を絶対パスに展開しておく
    prev, err := filepath.Abs(".")
    if err != nil {
        return // ERROR
    }
    defer os.Chdir(prev)

    // ディレクトリ移動
    os.Chdir("./repository/" + current_repo)

	// branch一覧取得
	branchs := git.Branch()
	for _, b := range strings.Split(string(branchs), "\n") {
		if  slashNum := strings.Index(b, ">"); slashNum != -1 {
			// アクティヴbranch
			continue
		} else if slashNum := strings.Index(b, "/"); slashNum == -1 {
			// 改行のみ
			continue
		} else {
			// branch格納
			b = b[slashNum+1:]
			var branch []string
			branch = append(branch, url)
			branch = append(branch, b)
			branch = append(branch, "2017")
			branchList = append(branchList, branch)
		}
	}

    // ディレクトリ移動
    os.Chdir(prev)

	// view情報
	h := Html{
		User:         user,
		Ip:           ip,
		Repository:   repository,
		BranchList:   branchList,
		Sidelink:     sidelink,
		Clone:        clone,
	}

	tmpl, _ := template.ParseFiles("./resources/view/branchs/index.html")
	tmpl.Execute(w, h)
}
