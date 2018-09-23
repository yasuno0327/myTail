package cmd

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var n int // 描写する行数

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myTail",
	Short: "myTail is go implementation of tail command",
	Run:   AnalysisArgument,
}

func init() {
	rootCmd.PersistentFlags().IntVar(&n, "n", 10, "Number of line that you want print.")
}

func AnalysisArgument(cmd *cobra.Command, args []string) {
	// goルーチンで実行すると関数実行のほうが早く終わってしまうので
	// 実行待ちチャネルを作る。
	wg := &sync.WaitGroup{}
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// go func内でAddしてしまうと処理が早すぎた際1回目の処理でWaitをthrowしてしまう。
	wg.Add(len(args))
	for i := range args {
		go func(filename string) {
			PrintFileN(n, filename, wd)
			wg.Done()
		}(args[i])
	}
	wg.Wait()
}

// 最後からn行分だけファイルの中身をプリントする
func PrintFileN(n int, filename string, wd string) {
	file, err := os.Open(wd + "/" + filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
	for i := 0; i < n && head < length; i++ {
		output := lines[head]
		fmt.Printf("%s\n", output)
		head++
	}
	fmt.Println("")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
