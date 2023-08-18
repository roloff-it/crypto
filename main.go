package main

import (
	"fmt"

	"github.com/roloff-it/crypto/decrypt"
	"github.com/roloff-it/crypto/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("This is very evil")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
	fmt.Println()
}
