package main

import (
	"log"

	"github.com/ivankuchin/sdwanctl/internal/cmd"
	configreader "github.com/ivankuchin/sdwanctl/internal/config-reader"
)

func SetLogFlags() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	SetLogFlags()

	_, err := configreader.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	cmd.Execute()
}
