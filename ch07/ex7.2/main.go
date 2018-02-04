package main

import (
	"bytes"
	"fmt"
	"io"
)

type countingWriter struct {
	io.Writer
	written *int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	bytesWritten, err := cw.Writer.Write(p)
	if err != nil {
		return bytesWritten, err
	}

	*cw.written = int64(bytesWritten)
	return bytesWritten, err
}

// CountingWriter takes an io.Writer and returns a new Writer that wraps the
// old Writer, and a pointer to an int64 that at any point in time holds the
// number of bytes written to the returned Writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var cntr int64
	return &countingWriter{w, &cntr}, &cntr
}

func main() {
	var bb bytes.Buffer
	cw, cntrPtr := CountingWriter(&bb)

	fmt.Fprint(cw, "this is a string")
	fmt.Println("count is: ", *cntrPtr) // count is: 16
}
