package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("-----------------------------------")
	write1()

	fmt.Println("-----------------------------------")
	write2()

	fmt.Println("-----------------------------------")
	write3()

	fmt.Println("-----------------------------------")
	write4()

	fmt.Println("-----------------------------------")
	write5()
}

/*
一次将全部内容写入文件
将数据写入文件的最快方法是使用该os.WriteFile()函数。它需要三个输入参数：
	我们要写入的文件的路径
	我们要写入文件的字节数据
	将创建的文件的权限[1]
	创建和关闭文件由函数本身完成，因此无需在写入前后创建或关闭文件。
如果您使用的是 1.16 之前的 Go 版本，您将WriteFile()在ioutil包中找到该功能。
*/
func write1() {
	if err := os.WriteFile("local/file1.txt", []byte("Hello GOSAMPLES!"), 0666); err != nil {
		log.Fatal(err)
	}
}

/*
将文本数据逐行写入文件
如果您将文件的行放在单独的变量、数组中，或者想在写入一行之前进行一些处理，则可以使用该func (*File) WriteString()方法逐行写入数据。
您需要做的就是创建一个文件，将字符串写入其中，最后关闭该文件。
*/
func write2() {
	var lines = []string{
		"Go",
		"is",
		"the",
		"best",
		"programming",
		"language",
		"in",
		"the",
		"world",
	}

	// create file
	f, err := os.Create("local/file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	for _, line := range lines {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

/*
将字节数据写入文件
	与逐行写入字符串一样，我们也可以使用该func (*File) Write()方法写入字节数据。
	或者func (*File) WriteAt()如果您想以给定的偏移量写入数据。
*/
func write3() {
	var bytes = []byte{
		0x47, // G
		0x4f, // O
		0x20, // <space>
		0x20, // <space>
		0x20, // <space>
		0x50, // P
		0x4c, // L
		0x45, // E
		0x53, // S
	}

	var additionalBytes = []byte{
		0x53, // S
		0x41, // A
		0x4d, // M
	}

	// create file
	f, err := os.Create("local/file3.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	// write bytes to the file
	_, err = f.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}

	// write additional bytes to the file, start at index 2
	_, err = f.WriteAt(additionalBytes, 2)
	if err != nil {
		log.Fatal(err)
	}
}

/*
将格式化后的字符串写入文件
除了File方法之外，我们还可以使用fmt.Fprintln()函数将数据写入文件。
此函数格式化其操作数，在它们之间添加空格，在末尾添加一个新行，并将输出写入 writer（第一个参数）。
它非常适合简单的行格式化或将 a 的字符串表示形式写入struct文件。
*/
func write4() {
	var lines = []string{
		"Go",
		"is",
		"the",
		"best",
		"programming",
		"language",
		"in",
		"the",
		"world",
	}

	// create file
	f, err := os.Create("local/file4.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	for _, line := range lines {
		_, err := fmt.Fprintln(f, "*", line, "*")
		if err != nil {
			log.Fatal(err)
		}
	}
}

/*
使用缓冲写入文件
如果您经常将少量数据写入文件，则会降低程序的性能。每次写入都是一个代价高昂的系统调用，如果您不需要立即更新文件，最好将这些小写入归为一个。

为此，我们可以使用bufio.Writer结构。
它的写入函数不会直接将数据保存到文件中，而是一直保存到下面的缓冲区已满（默认大小为 4096 字节）或Flush()调用该方法。
所以一定要Flush()在写入完成后调用，将剩余的数据保存到文件中。
*/
func write5() {
	var lines = []string{
		"Go",
		"is",
		"the",
		"best",
		"programming",
		"language",
		"in",
		"the",
		"world",
	}

	// create file
	f, err := os.Create("local/file5.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	// create new buffer
	buffer := bufio.NewWriter(f)

	for _, line := range lines {
		_, err := buffer.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	// flush buffered data to the file
	if err := buffer.Flush(); err != nil {
		log.Fatal(err)
	}
}
