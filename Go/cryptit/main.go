package main

import (
	"fmt"

	"github.com/itsaviral2609/AviralSingh-cryptit/decrypt"
	"github.com/itsaviral2609/AviralSingh-cryptit/encrypt"
)

func main() {
	EnCryptedStr := encrypt.Nimbus("Aviral")
	fmt.Println(EnCryptedStr)
	fmt.Println(decrypt.Nimbus(EnCryptedStr))
}
