package _fs

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func Read(path string) (data []byte, err error) {
	data, err = ioutil.ReadFile(path)
	if err != nil {
		return
	}
	return
}

func Remove(path string) (err error) {
	err = os.RemoveAll(path)
	return
}

func Exists(path string) (ok bool, err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = nil
		return
	} else if err != nil {
		return
	}
	ok = true
	return
}

func IsFile(path string) (ok bool, err error) {
	file, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
			return
		}
		return
	}
	ok = !file.IsDir()
	return
}

func IsDir(path string) (ok bool, err error) {
	file, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
			return
		}
		return
	}
	ok = file.IsDir()
	return
}

func Files(path string, suffix ...string) (files []string, err error) {

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	sep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() {
			var dirFiles []string
			dirFiles, err = Files(filepath.Join(path, sep, fi.Name()), suffix...)
			if err != nil {
				return
			}
			files = append(files, dirFiles...)
		} else {

			if hasSuffix(suffix, fi.Name()) {
				files = append(files, filepath.Join(path, sep, fi.Name()))
			}
		}
	}

	return
}

func hasSuffix(suffix []string, fileName string) bool {
	for _, v := range suffix {
		if strings.HasSuffix(fileName, v) {
			return true
		}
	}

	return false
}

func MakeDir(path string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		cmd := exec.Command("bash", "-c", fmt.Sprintf("mkdir -p %s", path))
		err = cmd.Run()
		if err != nil {
			return
		}
	}
	return
}

func Write(path string, content []byte) (err error) {
	err = ioutil.WriteFile(path, content, 0644)
	if err != nil {
		return
	}
	return
}

func RealName(path string) (string, error) {

	stat, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	switch stat.Mode().Type() {
	case fs.ModeSymlink:
		var name string
		name, err = os.Readlink(path)
		if err != nil {
			return "", err
		}
		return name, nil
	}

	return filepath.Base(path), nil
}
