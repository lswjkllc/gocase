package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("-----------------------------------")
	read1()

	fmt.Println("-----------------------------------")
	read2()

	fmt.Println("-----------------------------------")
	read3()

	fmt.Println("-----------------------------------")
	read4()
}

/*
读取整个文件
在 Go 中读取文本或二进制文件的最简单方法是使用 os 包中的 ReadFile() 函数。
此函数将文件的全部内容读到一个byte切片，因此在尝试读取大文件时应该注意 - 在这种情况下，您应该逐行或分块读取文件。
对于小文件，这种方式绰绰有余。

如果您使用的是 1.16 之前的 Go 版本，您将ReadFile()在`ioutil`[2]包中找到该功能。
*/
func read1() {
	content, err := os.ReadFile("local/file1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}

/*
逐行读取文件
要逐行读取文件，我们可以使用比较方便的bufio.Scanner结构。
它的构造函数 NewScanner() 接受一个打开的文件（记住在操作完成后关闭文件，例如通过 defer语句），并让您通过 Scan() 和 Text() 方法读取后续行。
使用 Err() 方法，您可以检查文件读取过程中遇到的错误。
*/
func read2() {
	// open file
	f, err := os.Open("local/file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		fmt.Printf("line: %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

/*
逐字的读取文件
逐字读取文件与逐行读取几乎相同。
您只需要将 Scanner 的 split 功能从**默认的 ScanLines() **函数更改为 ScanWords() 即可。
*/
func read3() {
	// open file
	f, err := os.Open("local/file3.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes) // bufio.ScanWords

	for scanner.Scan() {
		// do something with a word
		fmt.Println("char: " + scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

/*
分块读取文件
当你有一个非常大的文件或不想将整个文件存储在内存中时，您可以通过固定大小的块读取文件。
在这种情况下，您需要创建一个指定大小 chunkSize 的 byte 切片作为缓冲区，用于存储后续读取的字节。
使用 Read() 方法加载文件数据的下一个块。当发生 io.EOF 错误，指示文件结束，读取循环结束。
*/
func read4() {
	const chunkSize = 10

	// open file
	f, err := os.Open("local/file4.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	buf := make([]byte, chunkSize)

	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(buf[:n]))
	}
}
