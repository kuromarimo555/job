package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

var (
	filepath = flag.String("f", "", "filepash set")
)

func main() {
	flag.Parse()

	file, err := os.Open(*filepath)
	if err != nil {
		fmt.Println("ファイル無し")
		os.Exit(1)
	}
	scnr := bufio.NewScanner(transform.NewReader(file, japanese.ShiftJIS.NewDecoder()))

	for scnr.Scan() {
		fmt.Println(scnr.Text())
	}
	os.Exit(0)
}
