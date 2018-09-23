package cmd

import (
	"bufio"
	"fmt"
	"os"
)

// 最後からn行分だけファイルの中身をプリントする
func PrintFileN(n int, filename string, wd string) {
	file, err := os.Open(wd + "/" + filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// 最初にある程度容量を確保する(append高速化)
	lines := make([]string, 0, 1000)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	length := len(lines)
	if n > length {
		n = length
	}
	head := length - n // printする先頭行
	mutex.Lock()       // printの間だけ同期ロック
	fmt.Printf("==> %s <==\n", filename)
	for i := 0; i < n && head < length; i++ {
		output := lines[head]
		fmt.Printf("%s\n", output)
		head++
	}
	fmt.Println("")
	mutex.Unlock()
}
