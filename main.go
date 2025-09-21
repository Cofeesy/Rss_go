package main

import (
	"log"
	"os"
	"goinaction/search"
	_ "goinaction/matchers"
	"fmt"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	fmt.Println("a")
	search.Run("President")
}
