package utils

import (
	"bytes"
	"crypto/aes"
)

// AES ECB模式的加密解密
type AesTool struct {
	//128 192  256位的其中⼀个长度对应分别是 16 24  32字节长度
	Key       []byte
	BlockSize int
}

const AES_KEY = "kGx1ae1Sc0qCo88F"
const AES_BLOCK_SIZE = aes.BlockSize

func NewAesTool(key []byte, blockSize int) *AesTool {
	return &AesTool{Key: key, BlockSize: blockSize}
}

var DefaultAesTool = NewDefaultAesTool()

func NewDefaultAesTool() *AesTool {
	return NewAesTool([]byte(AES_KEY), AES_BLOCK_SIZE)
}

func (a *AesTool) padding(src []byte) []byte {
	//填充个数
	paddingCount := aes.BlockSize - len(src)%aes.BlockSize
	if paddingCount == aes.BlockSize { // 明文长度为16的倍数，则不填充
		return src
	} else {
		//填充数据
		return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
	}
}

// unpadding
func (a *AesTool) unPadding(src []byte) []byte {
	for i := len(src) - 1; i > 0; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
	return nil
}

func (a *AesTool) Encrypt(src []byte) ([]byte, error) {
	//key只能是 16 24 32长度
	block, err := aes.NewCipher([]byte(a.Key))
	if err != nil {
		return nil, err
	}
	//padding
	paddingSrc := a.padding(src)
	//返回加密结果
	encryptData := make([]byte, len(paddingSrc))
	//存储每次加密的数据
	tmpData := make([]byte, a.BlockSize)
	//分组分块加密
	for index := 0; index < len(paddingSrc); index += a.BlockSize {
		block.Encrypt(tmpData, paddingSrc[index:index+a.BlockSize])
		copy(encryptData[index:index+a.BlockSize], tmpData)
	}
	return encryptData, nil
}
func (a *AesTool) Decrypt(src []byte) ([]byte, error) {
	//key只能是 16 24 32长度
	block, err := aes.NewCipher([]byte(a.Key))
	if err != nil {
		return nil, err
	}
	//返回加密结果
	decryptData := make([]byte, len(src))
	//存储每次加密的数据
	tmpData := make([]byte, a.BlockSize)
	//分组分块加密
	for index := 0; index < len(src); index += a.BlockSize {
		block.Decrypt(tmpData, src[index:index+a.BlockSize])
		copy(decryptData[index:index+a.BlockSize], tmpData)
	}
	return a.unPadding(decryptData), nil
}
