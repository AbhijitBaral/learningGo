//~~~~~~~~~~~Rune type~~~~~~~~~~~~~~~~~~~~~~~~~~~
//Rune is a type that represents a unicode code point.
//It is essentially an alias for `int32`  

package main

import (
	"fmt"
)

func main() {
	// Example string with a mix of ASCII and non-ASCII characters
	str := "Hello, 世界!"

	// Convert the string to a slice of runes
	runes := []rune(str)

	// Print each rune as a character and its Unicode code point
	for _, r := range runes {
		fmt.Printf("Character: %c, Unicode: %U\n", r, r);
	}

	// Accessing individual runes
	fmt.Println("First rune:", string(runes[0])) // Output: H
	fmt.Println("Last rune:", string(runes[len(runes)-1])) // Output: !
}

