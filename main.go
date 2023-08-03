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
	
	Decrypt()
}

func Decrypt() {
	h := sha256.New()

	h.Write([]byte("kartikb385@gmail.com1234"))
	key := h.Sum([]byte(nil))

	file, err := os.OpenFile("nonce.txt", os.O_RDONLY, 0644)
	if err != nil {
		panic("Error in openeing file")
	}
	defer file.Close()
	r := bufio.NewReader(file)
	nonce :=make([]byte,12);
	n, err := r.Read(nonce)
	fmt.Println(n);
	if err != nil {
		panic("Error in Reading")

	}
	fmt.Println(nonce);
	block, err := aes.NewCipher(key)

	if err != nil {
		panic("error in generating gcm")
	}
	
	if err != nil {
		fmt.Printf("%v",err);
		panic("error in generating cipher")
		
	}
	file, err = os.OpenFile("cp.txt", os.O_RDONLY, 0644)

	if err != nil {
		panic("Error in openeing file")
	}
	defer file.Close()
	
	temp,err:= io.ReadAll(file)
	if err != nil {
		panic("Error in Reading")

	}
	cipher, err := cipher.NewGCM(block)
	if err != nil {
		panic("error in generating gcm")
	}

	fmt.Println(temp)
	cipherText,err := cipher.Open(nil, nonce, temp, nil)
	

	if err != nil {
		panic(err)
	}

	fmt.Println(string(cipherText))
}

func Encrypt(message string) {
    h := sha256.New()
    h.Write([]byte("kartikb385@gmail.com1234"))
    key := h.Sum([]byte(nil))

    // Generate a random nonce
    nonce := make([]byte, 12)
    if _, err := rand.Read(nonce); err != nil {
        panic(err)
    }
	fmt.Println(nonce);

	file, err := os.OpenFile("nonce.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    w := bufio.NewWriter(file)
    _, err = w.Write(nonce)
    if err != nil {
        panic(err)
    }

    err = w.Flush()
    if err != nil {
        panic(err)
    }



	fmt.Println(string(nonce));

    block, err := aes.NewCipher(key)
    if err != nil {
        panic(err)
    }

    cipher, err := cipher.NewGCM(block)
    if err != nil {
        panic(err)
    }

    cipherText := cipher.Seal(nil, nonce, []byte(message), nil)

	fmt.Println(cipherText)

    // Use defer to close the file
    file, err = os.OpenFile("cp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    w = bufio.NewWriter(file)
    _, err = w.Write(cipherText)
    if err != nil {
        panic(err)
    }

    err = w.Flush()
    if err != nil {
        panic(err)
    }
}
