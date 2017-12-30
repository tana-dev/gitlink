package files

import (
	"fmt"
	"html/template"
//	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
//	"bufio"
	"../../lib/appconfig"
	"../../lib/util"
)

const funcname = "files"

var (
	ip string
	user string
	repository map[string]string
	clone string
	sidelink map[string]string
	fileinfoList [][]string
	breadcrumbs map[string]string
	line string
	viewflg bool
	url string
	fpath string
)

type Html struct {
	Ip           string
	User         string
	Repository   map[string]string
	Clone        string
	Sidelink     map[string]string
	FileinfoList [][]string
	Breadcrumbs  map[string]string
	Line         string
	Viewflg      bool
//	Line         []string
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

	// breadcrumbs create
	breadcrumbs = map[string]string{}
	breadcrumbs = util.CreateBreadcrumbs(url, fpath, funcname)

	// ファイル存在チェック
	local_fpath := "./repository/" + strings.Replace(fpath, "/files/", "", 1)
	local_fpath = strings.TrimRight(local_fpath, "/")
	fi, err := os.Stat(local_fpath)
	if err != nil {
		fmt.Fprintf(w, "ファイル、もしくはディレクトが存在しません")
		return
	}

	//
	if fi.IsDir() {
		fpaths := util.Dirwalk(local_fpath)
		for _, fp := range fpaths {
			var fileinfo []string
			var dir string
			link := strings.Replace(fp, `\`, "/", -1)                 // 1.Windows
//			link = url + "/files" + strings.Replace(link, "/", "", 2) // 1.Windows
			link = url + "/files" + strings.Replace(link, "repository", "", 1) // 2.Linux
			name := filepath.Base(fp)
			f, _ := os.Stat(fp)

			// ファイルアイコン種類
			if f.IsDir() {
				dir = "fa-folder"
			} else {
				dir = "fa-file-o"
			}

			if err != nil {
				fmt.Fprintf(w, "ファイルの読み込みに失敗しました")
				return
			}
			updatetime_tmp := f.ModTime()
			updatetime := updatetime_tmp.Format("2006-01-02 15:04:05")

			fileinfo = append(fileinfo, link)
			fileinfo = append(fileinfo, name)
			fileinfo = append(fileinfo, updatetime)
			fileinfo = append(fileinfo, dir)
			fileinfoList = append(fileinfoList, fileinfo)
		}
		// sort.Sort(fileinfoList)

	} else {

		l, err := ioutil.ReadFile(local_fpath)
		if err != nil {
			// エラー処理
		}
		line = string(l[:])

		viewflg = true

	}

	// view情報
	h := Html{
		FileinfoList: fileinfoList,
		Breadcrumbs:  breadcrumbs,
		User:         user,
		Ip:           ip,
		Repository:   repository,
		Clone:        clone,
		Line:         line,
		Viewflg:      viewflg,
		Sidelink:     sidelink,
	}

	tmpl, _ := template.ParseFiles("./resources/view/files/index.html")
	tmpl.Execute(w, h)
}
