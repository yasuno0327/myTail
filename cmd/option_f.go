package cmd

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

func WatchFile(n int, filename string, wd string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer watcher.Close()

	done := make(chan bool)
	if err := watcher.Add(wd + "/" + filename); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op.String() == "WRITE" {
				PrintFileN(n, filename, wd)
			}
		case err := <-watcher.Errors:
			fmt.Println(err)
			os.Exit(1)
		}
	}
	<-done
}
