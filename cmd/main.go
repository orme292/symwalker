package main

import (
	"fmt"
	"os"

	"github.com/orme292/symwalker"
)

func main() {
	conf := swalker.SymConf{
		StartPath:      "/Users/aorme/",
		FollowSymlinks: true,
	}

	res, err := swalker.SymWalker(&conf)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, entry := range res {
		fmt.Printf("Path: %s\n", entry.Path)
	}
	os.Exit(0)
}
