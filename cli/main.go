package main

import (
	"fmt"
	"os"

	"github.com/certyclick_verify/core"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("Usage: certyclick_verify <file> <hash>")
		return
	}

	hashed, err := core.CalculateHash(args[1])
	if err != nil {
		fmt.Println("Error calculating file hash:", err)
		return
	}

	if fmt.Sprintf("%x", hashed) == args[2] {
		fmt.Println("The file hash matches the provided hash.")
	} else {
		fmt.Println("The file hash does not match the provided hash.")
		fmt.Println("Provided hash:", args[2])
		fmt.Println("Calculated hash:", fmt.Sprintf("%x", hashed))
	}
}
