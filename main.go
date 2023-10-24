package main

import (
	"log"

	"github.com/ivankuchin/phpipamsync/internal/cmd"
	"github.com/ivankuchin/phpipamsync/internal/config_reader"
)

func SetLogFlags() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

func main() {
	SetLogFlags()

	_, err := config_reader.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	cmd.Execute()
}
