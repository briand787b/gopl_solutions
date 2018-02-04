package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type wordCounter int

func (w *wordCounter) Write(p []byte) (int, error) {
	scnr := bufio.NewScanner(bytes.NewBuffer(p))
	scnr.Split(bufio.ScanWords)
	for scnr.Scan() {
		*w++
	}

	return len(p), nil
}

type lineCounter int

func (lc *lineCounter) Write(p []byte) (int, error) {
	scnr := bufio.NewScanner(bytes.NewBuffer(p))
	scnr.Split(bufio.ScanLines)
	for scnr.Scan() {
		*lc++
	}

	return len(p), nil
}

func main() {
	var wc wordCounter
	fmt.Fprintf(&wc, "this is four %s", "words")
	fmt.Println("words counted: ", wc)

	var lc lineCounter
	fmt.Fprintf(&lc, "this is \n two %s", "lines")
	fmt.Println("lines counted: ", lc)
}
