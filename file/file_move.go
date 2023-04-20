package file

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func Copy(from, to string) error {
	var err error
	f, err := os.Stat(from)
	if err != nil {
		return err
	}

	// 抽取公共的复制方法
	copyFn := func(fromFile string) error {
		// 复制文件的路径
		rel, err := filepath.Rel(from, fromFile)
		if err != nil {
			return err
		}
		toFile := filepath.Join(to, rel)

		// 创建复制文件目录
		if err = os.MkdirAll(filepath.Dir(toFile), os.ModePerm); err != nil {
			return err
		}

		// 读取源文件
		file, err := os.Open(from)
		if err != nil {
			return err
		}
		defer file.Close()

		bufReader := bufio.NewReader(file)
		// 创建复制文件用于保存
		outFile, err := os.Create(toFile)
		if err != nil {
			return err
		}
		defer outFile.Close()
		_, err = io.Copy(outFile, bufReader)
		return err
	}

	// 转绝对路径
	pwd, _ := os.Getwd()
	if !filepath.IsAbs(from) {
		from = filepath.Join(pwd, from)
	}
	if !filepath.IsAbs(to) {
		to = filepath.Join(pwd, to)
	}

	// 复制
	if f.IsDir() {
		return filepath.WalkDir(from, func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				// 创建目录
				if err = os.MkdirAll(path, os.ModePerm); err != nil {
					return err
				}
			} else {
				return copyFn(path)
			}
			return err
		})

	} else {
		return copyFn(from)
	}
}

// linux/mac 中直接调用 cp 命令
func directCopy(from, to string) error {
	command := exec.Command("cp", "-fr", from, to)
	err := command.Run()
	if err != nil {
		return err
	}
	return nil
}
