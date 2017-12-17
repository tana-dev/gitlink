package branchs

import (
	"fmt"
	"html/template"
//	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"../../lib/git"
	"../../lib/appconfig"
)

type Html struct {
	User         string
	Ip           string
	Repository map[string]string
	BranchList [][]string
	Sidelink     map[string]string
}

func Index(w http.ResponseWriter, r *http.Request) {

	var user string
	var ip string
	var branchList [][]string
	var url string
	var fpath string
	var repository map[string]string
	var sidelink map[string]string

	// ユーザー設定情報取得
	userConfig, err := appconfig.Parse("./config/user.json")
	if err != nil {
		fmt.Println("error ")
	}
	ip = userConfig.Host + ":"+ userConfig.Port
	url = userConfig.Protocol + "://"+ ip
	user = userConfig.Username

	// レポジトリー取得
	fpath = r.URL.Path
	fpath = strings.Replace(fpath, "/branchs/", "", 1)
//	fpath = strings.TrimRight(fpath, "/")
	slashNum := strings.Index(fpath, "/")
	current_repo := fpath[0:slashNum]

	// repositoryセット
	repository = map[string]string{}
	repos := dirwalk("./repository")
	for _, rp := range repos {
		//link := strings.Replace(fp, `\`, "/", -1)      // 1.Windows
		//link = url + strings.Replace(link, "/", "", 2) // 1.Windows
		link := url + "/files" + strings.Replace(rp, "repository", "", 1) + "/" // 2.Linux
		name := filepath.Base(rp)
		repository[link] = name
	}

	// sidelinkセット
	sidelink = map[string]string{}
	funk := []string{"files", "branchs"}
	for _, f := range funk {
		//link := strings.Replace(fp, `\`, "/", -1)      // 1.Windows
		//link = url + strings.Replace(link, "/", "", 2) // 1.Windows
		link := url + "/" + f + "/" + current_repo + "/" // 2.Linux
		sidelink[link] = f
	}

    // あらかじめ戻り先を絶対パスに展開しておく
    prev, err := filepath.Abs(".")
    if err != nil {
        return // ERROR
    }
    defer os.Chdir(prev)

    // ディレクトリ移動
    os.Chdir("./repository/fileshare")

    // ディレクトリ移動
    os.Chdir(prev)

    git.Branch()

	// branchList作成
	var branch []string
	branch = append(branch, url)
	branch = append(branch, "master")
	branch = append(branch, "2017")
	branchList = append(branchList, branch)

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
	fmt.Println("tanaka")

	// out, _ := exec.Command("git", "branch", "-r").Output()
	// fmt.Println(string(out))

	h := Html{
		User:         user,
		Ip:           ip,
		Repository:   repository,
		BranchList:   branchList,
		Sidelink:     sidelink,
	}

	tmpl, _ := template.ParseFiles("./resources/view/branchs/index.html")
	tmpl.Execute(w, h)
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
