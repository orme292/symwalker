package main

import (
	"fmt"
	"os"

	"github.com/orme292/symwalker"
)

func main() {
	conf := swalker.SymConf{
		StartPath:      "/Users/aorme/github/",
		FollowSymlinks: true,
	}

	res, err := swalker.SymWalker(&conf)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for i := range res {
		fmt.Printf("Path: %s\n", res[i].Path)
	}
	os.Exit(1)
}
