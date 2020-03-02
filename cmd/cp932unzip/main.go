package main

import (
	"fmt"
	"log"
	"os"

	. "github.com/moutend/cp932unzip/pkg/unzip"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	if len(args) < 2 {
		return nil
	}
	for _, path := range args[1:] {
		if err := Unzip(path); err != nil {
			return err
		}
		fmt.Printf("done\t%s\n", path)
	}
	return nil
}
