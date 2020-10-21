package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/sqweek/dialog"
	"io"
	"os"
	"path/filepath"
)

const (
	AppName    = "copyfiles"
	AppVersion = "v1.0"
	AppDate    = "2020-10-21"
	AppAuthor  = "niqingyang"
	AppEmail   = "niqy@qq.com"
)

func main() {

	var options struct {
		Dir     string `long:"dir" description:"the path of the copy files parent dir"`
		Files   string `long:"files" description:"the files for copy"`
		Dest    string `long:"dest" description:"the destination directory of the copy files"`
		Version bool   `short:"v" long:"version" description:"the application version info"`
	}

	p := flags.NewParser(&options, flags.Default)

	files, err := p.Parse()

	if err != nil {
		os.Exit(0)
	}

	if options.Version {
		fmt.Printf("%s %s.%s\n", AppName, AppVersion, AppDate)
		fmt.Printf("author: %s\n", AppAuthor)
		fmt.Printf("email: %s\n", AppEmail)

		os.Exit(0)
	}

	if options.Files != "" {
		files = append(files, options.Files)
	}

	if len(files) == 0 {
		dialog.Message("not found any copy files").Title("Error").Error()
		panic("not found any copy files")
	}

	if options.Dest == "" {
		dest, err := dialog.Directory().Browse()

		if err != nil {
			if err.Error() == dialog.ErrCancelled.Error() {
				fmt.Println("User Cancelled")
				os.Exit(1)
			} else {
				dialog.Message("%s", err).Title("Error").Error()
				panic(err)
			}
		}

		options.Dest = dest
	}

	fmt.Printf("output dir %s\n", options.Dest)

	var dst *os.File

	for _, file := range files {

		// 打开源文件
		src, err := os.Open(filepath.Join(options.Dir, file))
		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		defer src.Close()

		// 目录不存在则创建
		if !Exists(filepath.Dir(filepath.Join(options.Dest, file))) {
			os.MkdirAll(filepath.Dir(filepath.Join(options.Dest, file)), os.ModePerm)
		}

		// 创建目标文件
		dst, err = os.OpenFile(filepath.Join(options.Dest, file), os.O_WRONLY|os.O_CREATE, 0644)

		if err != nil {
			fmt.Printf("%s\n", err)
			continue
		}
		defer dst.Close()

		// 复制
		_, err = io.Copy(dst, src)

		if err != nil {
			fmt.Printf("copy file error %s\n", err)
		} else {
			fmt.Printf("copy file %s success\n", file)
		}
	}

}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false

	}
	return true

}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)

	if err != nil {
		return false
	}

	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}
