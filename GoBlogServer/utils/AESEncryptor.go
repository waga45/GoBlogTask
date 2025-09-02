package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type AESEncryptor struct {
	Key []byte
}

func NewAESEncryptor(key string) (*AESEncryptor, error) {
	// AES密钥长度必须是16, 24或32字节(对应AES-128, AES-192, AES-256)
	keyBytes := []byte(key)
	validLengths := map[int]bool{16: true, 24: true, 32: true}
	if !validLengths[len(keyBytes)] {
		return nil, errors.New("密钥长度必须是16, 24或32字节")
	}
	return &AESEncryptor{Key: keyBytes}, nil
}
func (a *AESEncryptor) EncryptCBC(plaintext []byte) (string, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return "", err
	}

	// 填充明文以满足块大小
	plaintext = pkcs7Pad(plaintext, aes.BlockSize)

	// 创建存储密文的缓冲区，需要额外空间存储IV
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// 生成随机初始化向量(IV)
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 创建CBC加密器
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// 返回Base64编码的密文
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptCBC 使用CBC模式解密
func (a *AESEncryptor) DecryptCBC(ciphertext string) ([]byte, error) {
	// 解码Base64密文
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}

	if len(decoded) < aes.BlockSize {
		return nil, errors.New("密文太短")
	}

	// 提取IV和实际密文
	iv := decoded[:aes.BlockSize]
	decoded = decoded[aes.BlockSize:]

	// 创建CBC解密器
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decoded, decoded)

	// 去除填充
	return pkcs7Unpad(decoded), nil
}

// EncryptGCM 使用GCM模式加密（推荐，更安全）
func (a *AESEncryptor) EncryptGCM(plaintext []byte) (string, error) {
	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return "", err
	}

	// 创建GCM模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 生成随机nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 加密并认证
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// 返回Base64编码的密文
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptGCM 使用GCM模式解密
func (a *AESEncryptor) DecryptGCM(ciphertext string) ([]byte, error) {
	// 解码Base64密文
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(a.Key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(decoded) < nonceSize {
		return nil, errors.New("密文太短")
	}

	// 提取nonce和实际密文
	nonce, ciphertextBytes := decoded[:nonceSize], decoded[nonceSize:]

	// 解密并验证
	return gcm.Open(nil, nonce, ciphertextBytes, nil)
}

// pkcs7Pad PKCS7填充
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7Unpad PKCS7去除填充
func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
