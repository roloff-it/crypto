package main

import (
	"fmt"

	"github.com/roloff-it/cypto/decrypt"
	"github.com/roloff-it/cypto/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("This is extra evil")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
	fmt.Println()
}
