package main

import (
	"ascii/internal/functions"
	"fmt"
	"os"
	"strings"
)

func main() {
	bannerfile := ""
	input := ""
	switch {
	case len(os.Args) > 3 || len(os.Args) < 2:
		fmt.Println("Check back on the length of arguments")
	case len(os.Args) == 3:
		bannerfile = os.Args[2]
		input = os.Args[1]
	default:
		bannerfile = "standard.txt"
		input = os.Args[1]
	}

	if !strings.HasSuffix(bannerfile, ".txt") {
		bannerfile += ".txt"
	}

	data, err := os.ReadFile(bannerfile)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		banner := string(data)
		fmt.Println(functions.AsciiConverter(input, banner))
	}

}
