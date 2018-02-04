package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func stringsNewReader(s string) io.Reader {
	return bytes.NewBufferString(s)
}

func main() {
	sr := stringsNewReader(`
		<html>
			<head></head>
			<body>
				<h1>this is an h1 tag</h1>
				<p>this a p tag</p>
			</body>
		</html>`,
	)

	doc, err := html.Parse(sr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
