package files

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
//	"bufio"
	"../../lib/appconfig"
)

type Html struct {
	FileinfoList [][]string
	Breadcrumbs  map[string]string
	User         string
	Ip           string
	Repository   map[string]string
//	Line         []string
	Line         string
	Viewflg      bool
	Sidelink     map[string]string
}

func Index(w http.ResponseWriter, r *http.Request) {

	var ip string
	var user string
	var url string
	var fileinfoList [][]string
	var breadcrumbs map[string]string
	var fpath string
//	var fname string
	var repository map[string]string
	var sidelink map[string]string
//	var line []string
	var line string
	var viewflg bool

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

	// pathを取るにはr.URL.Pathで受け取文末のスラッシュを削除
	//fpath = `\` + strings.Replace(r.URL.Path, "/", `\`, -1) // 1.Windows
	//fpath = strings.TrimRight(fpath, `\`)                   // 1.Windows
//	fname = filepath.Base(fpath)

	// ファイル存在チェック
	local_fpath := "./repository/" + strings.Replace(fpath, "/files/", "", 1)
	local_fpath = strings.TrimRight(local_fpath, "/")
	fi, err := os.Stat(local_fpath)
	if err != nil {
		fmt.Fprintf(w, "ファイル、もしくはディレクトが存在しません")
		return
	}

	// breadcrumbs create
	dirs_list := strings.Split(strings.TrimLeft(fpath, "/"), "/")
	breadcrumbs = map[string]string{}
	var indexs map[int]string
	indexs = map[int]string{}
	for i := 0; i < len(dirs_list); i++ {
		for l := 0; l <= i; l++ {
			if l == 0 {
				indexs[i] = "/" + dirs_list[l] + "/"
			} else {
				indexs[i] = indexs[i] + dirs_list[l] + "/"
			}
		}
		index := url + "/files" + indexs[i]
		breadcrumbs[index] = dirs_list[i]
	}

	if fi.IsDir() {
		fpaths := dirwalk(local_fpath)
		for _, fp := range fpaths {
			var fileinfo []string
			var dir string
			//link := strings.Replace(fp, `\`, "/", -1)      // 1.Windows
			//link = url + strings.Replace(link, "/", "", 2) // 1.Windows
			link := url + "/files" + strings.Replace(fp, "repository", "", 1) // 2.Linux
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

//		// ファイルオープン
//		fp, err := os.Open(local_fpath)
//		if err != nil {
//			fmt.Fprintf(w, "ファイルの読み込みに失敗しました")
//		}
//		defer fp.Close()
//
//		scanner := bufio.NewScanner(fp)
//
//		// ファイルデータ格納
//		for scanner.Scan() {
//			s := strings.TrimRight(scanner.Text(), "\n")
//			line = append(line, s)
//		}
//
//		// 確認
//		if err = scanner.Err(); err != nil {
//			fmt.Fprintf(w, "ファイルの読み込みに失敗しました")
//		}

		l, err := ioutil.ReadFile(local_fpath)
		if err != nil {
			// エラー処理
		}
		line = string(l[:])

		viewflg = true

	}

	h := Html{
		FileinfoList: fileinfoList,
		Breadcrumbs:  breadcrumbs,
		User:         user,
		Ip:           ip,
		Repository:   repository,
		Line:         line,
		Viewflg:      viewflg,
		Sidelink:     sidelink,
	}

	tmpl, _ := template.ParseFiles("./resources/view/files/index.html")
	tmpl.Execute(w, h)
}

func readfile(srcpath string) []byte {

	src, err := os.Open(srcpath)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	contents, _ := ioutil.ReadAll(src)

	return contents
}

func copyfile(srcpath string, dstpath string) {

	src, err := os.Open(srcpath)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(dstpath)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}
}

func createContentType(ext string) string {

	var ctype string

	// w.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(name)))

	switch ext {
	case ".txt":
		ctype = "text/plain"
	case ".csv":
		ctype = "text/csv" // CSVファイル
	case ".html":
		ctype = "text/html" // HTMLファイル
	case ".css":
		ctype = "text/css" // CSSファイル
	case ".js":
		ctype = "text/javascript" // JavaScriptファイル
	case ".exe":
		ctype = "application/octet-stream" // EXEファイルなどの実行ファイル
	case ".pdf":
		ctype = "application/pdf" // PDFファイル
	case ".xlsx":
		// ctype = "application/vnd.ms-excel" // EXCELファイル
		ctype = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" // EXCELファイル
	case ".ppt":
		ctype = "application/vnd.ms-powerpoint" // PowerPointファイル
	case ".docx":
		ctype = "application/msword" // WORDファイル
	case ".jpeg", ".jpg":
		ctype = "image/jpeg" // JPEGファイル(.jpg, .jpeg)
	case ".png":
		ctype = "image/png" // PNGファイル
	case ".gif":
		ctype = "image/gif" // GIFファイル
	case ".bmp":
		ctype = "image/bmp" // Bitmapファイル
	case ".zip":
		ctype = "application/zip" // Zipファイル
	case ".lzh":
		ctype = "application/x-lzh" // LZHファイル
	case ".tar":
		ctype = "application/x-tar" // tarファイル/tar&gzipファイル
	case ".mp3":
		ctype = "audio/mpeg" // MP3ファイル
	case ".mp4":
		ctype = "audio/mp4" // MP4ファイル
	case ".mpeg":
		ctype = "video/mpeg" // MPEGファイル（動画）
	default:
		ctype = "text/plain"
	}

	return ctype
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
