package main

import (
	"flag"
	"fmt"
)
var numTest int
func init() {
	flag.IntVar(&numTest, "num",0,"numInput")
}
func main() {
	flag.Parse()
	fmt.Println(&numTest)
	fmt.Println(numTest)
}