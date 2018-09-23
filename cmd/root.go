// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "myTail",
	Short: "myTail is go implementation of tail command",
	Run:   AnalysisArgument,
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
			PrintFile(filename, wd)
			wg.Done()
		}(args[i])
	}
	wg.Wait()
}

// ファイルの中身をプリントする
func PrintFile(filename string, wd string) {
	bytes, err := ioutil.ReadFile(wd + "/" + filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(bytes), "\n")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
