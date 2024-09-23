解密decrypt

package main

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"path/filepath"
	"os"
)

func decryptPath(path string, gcm cipher.AEAD) error {
	return filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", filePath, err)
			return err
		}
		// skip if directory or not .enc file
		if !info.IsDir() && filepath.Ext(filePath) == ".enc" {
			// decrypt the file
			fmt.Println("Decrypting " + filePath + "...")
			// read file contents
			encrypted, err := os.ReadFile(filePath)
			if err == nil {
				// Decrypt bytes
				nonce := encrypted[:gcm.NonceSize()]
				encrypted = encrypted[gcm.NonceSize():]
				original, err := gcm.Open(nil, nonce, encrypted, nil)		
				if err != nil {
					fmt.Printf("Error decrypting file %q: %v\n", filePath, err)
					return nil
				}

				// write decrypted contents
				decryptedPath := filePath[:len(filePath)-4] // remove .enc extension
				err = os.WriteFile(decryptedPath, original, 0666)
				if err == nil {
					os.Remove(filePath) // delete the encrypted file
					fmt.Printf("Successfully decrypted and removed: %s\n", filePath)
				} else {
					fmt.Printf("Error writing decrypted contents to %q: %v\n", decryptedPath, err)
				}
			} else {
				fmt.Printf("Error reading file contents from %q: %v\n", filePath, err)
			}
		}
		return nil
	})
}

func main() {
	fmt.Println("Hello^_^")
	fmt.Print("Key: ")
	var key string
	fmt.Scanln(&key)
	
	// Initialize AES in GCM mode
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic("error while setting up aes")
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic("error while setting up gcm")
	}

	// 定義兩個路徑
	path1 := `C:\Users\aaaa\Downloads`
	path2 := `C:\Users\aaaa\Desktop`

	// 解密第一個路徑
	fmt.Println("Decrypting files in", path1)
	if err := decryptPath(path1, gcm); err != nil {
		fmt.Printf("Error decrypting %s: %v\n", path1, err)
	}

	// 解密第二個路徑
	fmt.Println("Decrypting files in", path2)
	if err := decryptPath(path2, gcm); err != nil {
		fmt.Printf("Error decrypting %s: %v\n", path2, err)
	}

	fmt.Println("Decryption process completed.")
}

