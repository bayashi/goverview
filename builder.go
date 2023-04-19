package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	pt "github.com/bayashi/go-proptree"
	"github.com/bayashi/goverview/fileinfo"
	ignore "github.com/sabhiram/go-gitignore"
)

func fromLocal(o *options) (*pt.N, error) {
	path, err := validateDirPath(o.path)
	if err != nil {
		return nil, err
	}

	return buildTree(o, path)
}

func validateDirPath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	d, err := os.Stat(absPath)
	if err != nil || !d.IsDir() {
		return "", err
	}

	if !d.IsDir() {
		return "", fmt.Errorf("not directory: %s", path)
	}

	return absPath, nil
}

type walkerArgs struct {
	path        string
	f           os.FileInfo
	relPath     string
	parentN     *pt.N
	s           *map[string]*pt.N
	currentPath *string
	o           *options
}

func buildTree(o *options, dirPath string) (*pt.N, error) {
	gitignoreFilePath := strings.Join([]string{dirPath, ".gitignore"}, string(os.PathSeparator))
	gitignore, _ := ignore.CompileIgnoreFile(gitignoreFilePath)

	dirPathLen := len(dirPath)
	rootN := pt.Node(filepath.Base(dirPath) + "/")
	s := map[string]*pt.N{"": rootN}
	currentPath := ""

	walkErr := filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
		if len(path) == dirPathLen {
			return nil // skip rootN
		}
		if (gitignore != nil && gitignore.MatchesPath(path)) || isSkipPath(o, path) {
			return nil
		}

		arg := &walkerArgs{
			path:        path,
			f:           f,
			relPath:     path[dirPathLen+1:],
			parentN:     rootN,
			s:           &s,
			currentPath: &currentPath,
			o:           o,
		}
		if err := walkProcess(arg); err != nil {
			return err
		}

		return nil
	})

	return rootN, walkErr
}

const (
	pathIDSeparator = "__"
)

func walkProcess(arg *walkerArgs) error {
	ss := *arg.s
	cPath := *arg.currentPath
	elements := strings.Split(arg.relPath, string(os.PathSeparator))
	for i, el := range elements {
		cPath = cPath + el + pathIDSeparator
		if i == len(elements)-1 {
			var n *pt.N
			var err error
			if arg.f.IsDir() {
				n = pt.Node(el + "/")
			} else {
				n, err = getFileInfo(pt.Node(el), arg)
				if err != nil {
					return err
				}
			}
			arg.parentN.Append(n)
			ss[cPath] = n
			continue
		}

		if _, isExists := ss[cPath]; !isExists {
			ss[cPath] = pt.Node(cPath)
		}
		arg.parentN = ss[cPath]
	}

	return nil
}

func getFileInfo(n *pt.N, arg *walkerArgs) (*pt.N, error) {
	if filepath.Ext(arg.path) == ".go" {
		fi, err := fileinfo.GoInfo(arg.path, arg.o.showAll)
		if err != nil {
			return nil, err
		}
		n.Icon(fi.Icon).
			Tag(fi.Tag).
			Descriptions(fi.Descriptions)
	} else if filepath.Base(arg.path) == "go.mod" {
		fi, err := fileinfo.GoModInfo(arg.path, arg.o.showAll)
		if err != nil {
			return nil, err
		}
		n.Tag(fi.Tag)
	} else if filepath.Base(arg.path) == "LICENSE" {
		fi, err := fileinfo.LicenseInfo(arg.path, arg.o.showAll)
		if err != nil {
			return nil, err
		}
		n.Tag(fi.Tag)
	}

	return n, nil
}

func isSkipPath(o *options, path string) bool {
	if strings.Contains(path, "/.git") &&
		!strings.Contains(path, "/.gitignore") && !strings.Contains(path, "/.github") {
		return true // skip
	}

	for _, i := range o.ignore {
		if strings.Contains(path, i) {
			return true // skip
		}
	}

	return false
}
