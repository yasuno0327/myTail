package cmd

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

var (
	n     int        // 描写する行数
	mutex sync.Mutex // print用ロック
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myTail",
	Short: "myTail is go implementation of tail command",
	Run:   AnalyzeArgument,
}

func init() {
	rootCmd.PersistentFlags().IntVar(&n, "n", 10, "Number of line that you want print.")
}

func AnalyzeArgument(cmd *cobra.Command, args []string) {
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
