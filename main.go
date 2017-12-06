package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	str := os.Args[1]
	if str == "" {
		log.Fatalln("must str parameter: ex) bc 'hogehoge' ")
	}

	b := []byte(str)
	fmt.Printf("%d\n", len(b))
}
