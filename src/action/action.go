package action

import (
	"fmt"
	// "os/exec"
	// "../git"
)

func Regist() {

	var url string
	var currentDir string

	url = "http://10.27.145.100:8080/"
	url = "http://192.168.33.22:8080/"

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
    action.Regist()
	fmt.Println("tanaka")

	// out, _ := exec.Command("git", "branch", "-r").Output()
	// fmt.Println(string(out))

	// tmpl := template.Must(template.ParseFiles("./view/index.html"))
	// tmpl.Execute(w, h)
}

func Diff() {
	fmt.Println("diff")
}
