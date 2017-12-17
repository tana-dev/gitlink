package action

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

// var url = "http://10.27.145.100:12000/"
var url = "http://192.168.33.22:12000/"
var user = "tanaka-shu"

func Regist() {

	var url string
	var currentDir string

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

	// tmpl := template.Must(template.ParseFiles("./view/index.html"))
	// tmpl.Execute(w, h)
}

func Files() {

    var url string
    var fileinfoList [][]string
    var breadcrumbs map[string]string
    var fpath string
    var fname string
    var bookmark map[string]string

    // url = "http://" + ip + "/"

    // pathを取るにはr.URL.Pathで受け取文末のスラッシュを削除
    // fpath = `\` + strings.Replace(r.URL.Path, "/", `\`, -1) // 1.Windows
    // fpath = strings.TrimRight(fpath, `\`) // 1.Windows
    fpath = strings.TrimRight(fpath, "/") // 2. Linux
    fname = filepath.Base(fpath)

    // ファイル存在チェック
    fi, err := os.Stat(fpath)
    if err != nil {
        fmt.Fprintf(w, "ファイル、もしくはディレクトが存在しません")
        return
    }

    // breadcrumbs create
    dirs_list := strings.Split(strings.TrimLeft(fpath1, "/"), "/")
    breadcrumbs = map[string]string{}
    var indexs map[int]string
    indexs = map[int]string{}
    for i := 0; i < len(dirs_list); i++ {
        for l := 0; l <= i; l++ {
            if l == 0 {
                indexs[i] = dirs_list[l] + "/"
            } else {
                indexs[i] = indexs[i] + dirs_list[l] + "/"
            }
        }
        index := url + indexs[i]
        breadcrumbs[index] = dirs_list[i]
    }

    if fi.IsDir() {
        fpaths := dirwalk(fpath)
        for _, fp := range fpaths {
            var fileinfo []string
            var dir string
            // link := strings.Replace(fp, `\`, "/", -1)       // 2.Windows
            // link = url + strings.Replace(link, "/", "", 2)  // 2.Windows
            link := url + strings.Replace(fp, "/", "", 1) // 2.Linux
            name := filepath.Base(fp)
            f, _ := os.Stat(fp)
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

    } else {
        ext := fname[strings.LastIndex(fname, "."):]
        out := readfile(fpath)
        ctype := createContentType(ext)
        w.Header().Set("Content-Disposition", "attachment; filename="+fname)
        w.Header().Set("Content-Type", ctype)
        // w.Header().Set("Content-Length", string(len(out)))
        w.Write(out)
        return
    }

    h := Html{
        FileinfoList: fileinfoList,
        Breadcrumbs:  breadcrumbs,
    }

    // funcs := template.FuncMap{"add": add}
    // tmpl := template.Must(template.New("./view/index.html").Funcs(funcs).ParseFiles("./view/index.html"))
    // tmpl.Execute(w, h)

    templ_file, err := Asset("../static/view/files/index.html")
    tmpl, _ := template.New("tmpl").Parse(string(templ_file))
    tmpl.Execute(w, h)

	fmt.Println("files")
}

func Diff() {
	fmt.Println("diff")
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
        if 0 != strings.Index(file.Name(), ".") && 0 != strings.Index(file.Name(), "~$") {

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

