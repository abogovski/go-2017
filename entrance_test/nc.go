package main

import (
    "fmt"
    "os"
    "io"
    "flag"
    "strings"
)

func Check(err error) {
    if err != nil {	
	fmt.Fprintf(os.Stderr, err.Error())
        os.Exit(1)
    }
}

func CountNewlines(filename string, buf_size uint) uint64 {
    var newline_cnt uint64 = 0
    buf := make([]byte, buf_size)

    f, err := os.Open(filename)
    Check(err)

    for {
        n_read, err := f.Read(buf)
        newline_cnt = newline_cnt + uint64(strings.Count(string(buf[:n_read]), "\n"))
	if err != nil {
            f.Close()
	    if err == io.EOF {
	        return newline_cnt
	    }
	    Check(err)
	}
    }
}

func Usage() {
    fmt.Fprintf(os.Stderr, "\n")
    fmt.Fprintf(os.Stderr, "Usage:\n  %s [flags] filename\n", os.Args[0])
    fmt.Fprintf(os.Stderr, "\n")
    fmt.Fprintf(os.Stderr, "Flags:\n")
    flag.PrintDefaults()
}

func main() {
    flag.Usage = Usage
    buf_size := flag.Uint("b", 4096 * 16, "buffer size") // default value ~ 16 disc sectors
    flag.Parse()

    if len(flag.Args()) != 1 {
        Usage()
	os.Exit(1)
    } else {
        fmt.Println(CountNewlines(flag.Arg(0), *buf_size))
    }
}
