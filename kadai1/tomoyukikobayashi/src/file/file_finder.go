/*
Package file provides methods and structs
that handle file and directory operations.
Almost all of them just wrap primitive methods
defined in golang native packages.
*/
package file

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Find returns all filepaths in directory specified as dir.
// If you want to filter the paths by file extensions, use exts.
func Find(dir string, exts []string) ([]string, error) {
	paths, err := find(dir)
	if err != nil {
		return nil, err
	}

	return paths.filter(exts), nil
}

// find dirで指定したディレクトリ配下のファイルパス一覧を返却
func find(dir string) (paths, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("could not open dir %s", dir)
	}

	paths := paths{}
	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		// FileInfoがディレクトリなら再帰処理
		if file.IsDir() {
			dsc, err := find(path)
			if err != nil {
				return nil, fmt.Errorf("could not open dir %s", path)
			}
			// ...を付けるとslice同士連結できるよう
			paths = append(paths, dsc...)
			continue
		}
		paths = append(paths, path)
	}

	return paths, nil
}

// paths ディレクトリパスの集合に対する操作を提供する
type paths []string

// filter extsで指定した拡張子にマッチする結果を絞り込む
// filter(条件1).filter(条件2) のようにして順に絞り込んでもよい
func (p paths) filter(exts []string) paths {
	paths := paths{}
	for _, v := range p {
		if matches(v, exts) {
			paths = append(paths, v)
		}
	}
	return paths
}

// matches ignore caseでpathがextsで指定した拡張子にマッチするかを検証する
func matches(path string, exts []string) bool {
	lpath := strings.ToLower(path)
	for _, v := range exts {
		if strings.HasSuffix(lpath, "."+strings.ToLower(v)) {
			return true
		}
	}
	return false
}
