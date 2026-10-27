package main

import (
	"fmt"
	"os"
)

func main() {
	// will be completed along the way
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(bookworms)
}
