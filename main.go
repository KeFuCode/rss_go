package main

import (
	"log"	
	"os"
	"go_code/goinaction/chapter2/sample2/search"
	_ "go_code/goinaction/chapter2/sample2/matchers"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}