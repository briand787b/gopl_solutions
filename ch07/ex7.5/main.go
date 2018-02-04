package main

import (
	"fmt"
	"io"
	"strings"
)

// EXERCISE: Create my own implementation of the LimitReader function from package "io"

type limitedReader struct {
	R io.Reader
	N int // bytes left to read
}

func (r *limitedReader) Read(p []byte) (n int, err error) {
	if r.N <= 0 {
		return 0, io.EOF
	}
	if len(p) > r.N {
		p = p[:r.N]
	}
	n, err = r.R.Read(p)
	r.N -= n
	return
}

func limitReader(r io.Reader, n int) io.Reader { return &limitedReader{r, n} }

func main() {
	sr := strings.NewReader("this is longer than 10 characters")
	lr := limitReader(sr, 10)

	bSlice := make([]byte, 121)
	n, err := lr.Read(bSlice)
	fmt.Println("error is: ", err)       // nil
	fmt.Println("bytes written is: ", n) // 10
	fmt.Println("bSlice: ", bSlice)      // [116, 104, ... 0] (length is 121 bytes)
}
