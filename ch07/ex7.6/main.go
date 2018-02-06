package main

import (
	"flag"
	"fmt"
	"github.com/briand787b/gopl_solutions/ch07/ex7.6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
