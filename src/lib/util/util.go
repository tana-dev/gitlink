package util

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func CreateSidevar(url string,current_repo string) map[string]string {

	sidelink := map[string]string{}
	funk := []string{"files", "branchs"}

	for _, f := range funk {
		link := url + "/" + f + "/" + current_repo + "/"
		sidelink[link] = f
	}

	return sidelink
}

func CreateBreadcrumbs(url string, fpath string, funcname string) map[string]string {

	//var indexs map[int]string
	//var breadcrumbs map[string]string
	breadcrumbs := map[string]string{}
	indexs := map[int]string{}

	dirs_list := strings.Split(strings.TrimLeft(fpath, "/"), "/")

	for i := 0; i < len(dirs_list); i++ {
		for l := 0; l <= i; l++ {
			if l == 0 {
				indexs[i] = "/" + dirs_list[l] + "/"
			} else {
				indexs[i] = indexs[i] + dirs_list[l] + "/"
			}
		}
		index := url + "/" + funcname + indexs[i]
		breadcrumbs[index] = dirs_list[i]
	}

	return breadcrumbs
}

func CreateRepository(url string) map[string]string {

	repository := map[string]string{}
	repos := Dirwalk("repository")

	for _, rp := range repos {
		var link string
		var name string
		link = strings.Replace(rp, `\`, "/", -1)                                 // 1.Windows
//		link = url + strings.Replace(link, "/", "", 2) // 1.Windows
		link = url + "/files" + strings.Replace(link, "repository", "", 1) + "/" // 2.Linux
		name = filepath.Base(rp)
		repository[link] = name
	}

	return repository
}

func Readfile(srcpath string) []byte {

	src, err := os.Open(srcpath)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	contents, _ := ioutil.ReadAll(src)

	return contents
}

func Copyfile(srcpath string, dstpath string) {

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

func Dirwalk(dir string) []string {

	//
	var paths []string
	var dpaths []string
	var fpaths []string

	//
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	//
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

	//
	if nil == dpaths && nil != fpaths {
		paths = fpaths
	} else if nil != dpaths && nil == fpaths {
		paths = dpaths
	} else {
		paths = append(dpaths, fpaths...)
	}

	return paths
}
