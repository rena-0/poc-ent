// +build ignore

package main

import (
	"log"

	// Import the Ent code generation tool
	_ "entgo.io/ent/cmd/entc" // This is required for generating the code
	"github.com/rena-0/poc-ent/ent"
)

func main() {
	if err := ent.Generate(); err != nil {
		log.Fatalf("running ent generate: %v", err)
	}
}