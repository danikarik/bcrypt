package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var (
	cost  = flag.Int("cost", bcrypt.DefaultCost, "Cost for bcrypt hasher")
	input = flag.String("input", "test", "Message to be hashed")
)

func main() {
	flag.Parse()
	fmt.Printf("message: %s, cost: %d\n", *input, *cost)
	hash, err := bcrypt.GenerateFromPassword([]byte(*input), *cost)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(hash))
}
