package service

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.design/x/clipboard"
)

func GeneratePassword(lenPass int, symbols bool) string {
	var pass string
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	lettersWithSym := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@$%^&*()_+1234567890")

	if symbols {
		letters = lettersWithSym
	}
	max := big.NewInt(int64(len(letters)))

	for i := 0; i < lenPass; i++ {
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err)
		}
		randomIndex := n.Int64()
		pass = pass + string(letters[randomIndex])
	}

	err := clipboard.Init()
	if err != nil {
		fmt.Print("error while coping")
	}
	clipboard.Write(clipboard.FmtText, []byte(pass))

	fmt.Printf("Copied text in serevice: \"%s\"\n", pass)

	return pass
}
