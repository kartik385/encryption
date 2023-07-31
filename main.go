package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"os"
)

func main() {
	h := sha256.New()

	h.Write([]byte("kartikb385@gmail.com1234"))
	key := h.Sum([]byte(nil))
	fmt.Println(h.Size())
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Printf("%v",err);
		panic("error in generating cipher")
		
	}
	cipher, err := cipher.NewGCM(block)
	if err != nil {
		panic("error in generating gcm")
	}
	cipherText := cipher.Seal(nil, nonce, []byte("Hello World"), nil)
	file, err := os.OpenFile("cp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic("Error in openeing file")
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	_, err = w.Write(cipherText)
	if err != nil {
		panic("Error in writing")

	}
	err = w.Flush()
	if err != nil {
		panic(err)
	}

}
