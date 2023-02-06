package main

import (
	"fmt"

	"github.com/itsaviral2609/AviralSingh-cryptit/encrypt"
	"github.com/pborman/uuid"
)

func main() {
	uuid := uuid.NewRandom()
	fmt.Println(uuid)

	EnCryptedStr := encrypt.Nimbus("Aviral")
	fmt.Println(EnCryptedStr)
}
