package main

import (
	"fmt"
	"os"
)

func main() {
	// will be completed along the way
	bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		// _, _ ignores returns by Fprintf
		// go functions inherently returns errors along with the expected return values, so we need to handle them
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(bookworms)
}
