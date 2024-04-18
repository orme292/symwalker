package main

import (
	"fmt"
	"os"

	sw "github.com/orme292/symwalker"
)

func main() {

	conf := sw.NewSymConf(
		sw.WithStartPath("/Users/andrew"),
		sw.WithFollowedSymLinks(),
		sw.WithLogging(),
	)

	results, err := sw.SymWalker(conf)
	if err != nil {
		fmt.Printf("Error occurred: %s", err.Error())
		os.Exit(1)
	}

	for _, dir := range results.Dirs {
		fmt.Printf("Dir: %s\n", dir.Path)
	}

	for _, file := range results.Files {
		fmt.Printf("File: %s\n", file.Path)
	}

	os.Exit(0)
}
