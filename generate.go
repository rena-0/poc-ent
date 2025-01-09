package main

import (
	"log"

	_ "github.com/ent/ent/cmd/entc" // Import ent codegen tool
	"github.com/rena-0/poc-ent/ent"  // Your ent package
)

func main() {
	if err := ent.Generate(); err != nil {
		log.Fatalf("running ent generate: %v", err)
	}
}