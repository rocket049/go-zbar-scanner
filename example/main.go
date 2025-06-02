package main

import (
	"flag"

	"github.com/rocket049/go-zbar-scanner"
)

func main() {
	flag.Parse()
	ret, err := goZbarScanner.ScanFile(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	println("Result:", ret)
}
