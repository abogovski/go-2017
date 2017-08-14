package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func countNewlines(filename string, bufSize uint) uint64 {
	var newlineCnt uint64
	buf := make([]byte, bufSize)

	f, err := os.Open(filename)
	check(err)

	for {
		nRead, err := f.Read(buf)
		newlineCnt = newlineCnt + uint64(strings.Count(string(buf[:nRead]), "\n"))
		if err != nil {
			f.Close()
			if err == io.EOF {
				return newlineCnt
			}
			check(err)
		}
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Usage:\n  %s [flags] filename\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	bufSize := flag.Uint("b", 4096*16, "Buffer Size") // default value ~ 16 disc sectors
	flag.Parse()

	if len(flag.Args()) != 1 {
		usage()
		os.Exit(1)
	} else {
		fmt.Println(countNewlines(flag.Arg(0), *bufSize))
	}
}
