package main

import (
	"log"

	cmd "github.com/hamza-007/ret3iac/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
