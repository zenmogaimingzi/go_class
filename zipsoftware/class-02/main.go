package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// 获取命令行参数（忽略前两个参数，因为一个是程序名 另一个是压缩文件目标路径）
	args := os.Args[2:]
	args1 := os.Args[1:]
	// 将 args 数组放到 files 切片中
	files := append([]string{}, args...)
	// 打印文件名列表
	fmt.Println("文件列表:", files)
	now := time.Now()
	// 格式化为 年-月-日 的形式
	formattedDate := now.Format("2006-01-02")
	//fmt.Println("格式化日期:", formattedDate)
	pathNames := append([]string{}, args1...)
	var pathName = pathNames[0]
	output := pathName + formattedDate + "-xxx.zip"
	// 判断 files 是否为空
	if len(files) == 0 {
		fmt.Println("没有提供任何文件参数！")
	} else {
		if err := ZipFiles(output, files); err != nil {
			panic(err)
		}
		fmt.Println("Zipped File:", output)
	}

}

// ZipFiles compresses one or many files into a single zip archive file.
// Param 1: filename is the output zip file's name.
// Param 2: files is a list of files to add to the zip.
func ZipFiles(filename string, files []string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filepathStr string) error {

	fileToZip, err := os.Open(filepathStr)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	//
	paths, fileName := filepath.Split(filepathStr)
	//获取路径中的目录及文件名 E:\data\test.txt
	fmt.Println(paths, fileName)
	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = fileName

	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate

	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}
